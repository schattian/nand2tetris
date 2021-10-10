package scanner

import (
	"reflect"
	"testing"

	"github.com/schattian/nand2tetris/compiler/token"
)

func TestScanner_Scan(t *testing.T) {
	type fields struct {
		src    []byte
		char   rune
		offset int
	}
	tests := []struct {
		name    string
		fields  fields
		wantTok token.Token
		wantLit string
	}{
		{
			name: "multiline comment then, eof",
			fields: fields{
				src: []byte(`// foo bar
		`),
			},
			wantTok: token.EOF,
			wantLit: string(eof),
		},
		{
			name: "multiline comment then var",
			fields: fields{
				src: []byte(`// foo bar
						// bar
		var a int = 1;`),
			},
			wantTok: token.VAR,
			wantLit: "var",
		},
		{
			name: "commented var",
			fields: fields{
				src: []byte(`// var a int = 1;`),
			},
			wantTok: token.EOF,
			wantLit: string(eof),
		},
		{
			name: "div with space",
			fields: fields{
				src: []byte(`/ 3;`),
			},
			wantTok: token.DIV,
			wantLit: "/",
		},
		{
			name: "int operation",
			fields: fields{
				src: []byte(`333 + 1;`),
			},
			wantTok: token.INTEGER_CONST,
			wantLit: "333",
		},
		{
			name: "div",
			fields: fields{
				src: []byte(`/3;`),
			},
			wantTok: token.DIV,
			wantLit: "/",
		},
		{
			name: "var",
			fields: fields{
				src: []byte(`var a int = 1;`),
			},
			wantTok: token.VAR,
			wantLit: "var",
		},
		{
			name: "varname def",
			fields: fields{
				src:    []byte(`var a int = 1;`),
				offset: 4,
			},
			wantTok: token.IDENT,
			wantLit: "a",
		},
		{
			name: "varname def a_1_a",
			fields: fields{
				src:    []byte(`var a_1_a int = 1;`),
				offset: 4,
			},
			wantTok: token.IDENT,
			wantLit: "a_1_a",
		},
		{
			name: "int type",
			fields: fields{
				src:    []byte(`var a int = 1;`),
				offset: 6,
			},
			wantTok: token.INT,
			wantLit: "int",
		},
		{
			name: "eq assign",
			fields: fields{
				src: []byte(`= 1;`),
			},
			wantTok: token.EQ,
			wantLit: "=",
		},
		{
			name: "integer const",
			fields: fields{
				src:    []byte(`= 1;`),
				offset: 2,
			},
			wantTok: token.INTEGER_CONST,
			wantLit: "1",
		},
		{
			name: ";",
			fields: fields{
				src:    []byte(`= 1;`),
				offset: 3,
			},
			wantTok: token.SEMICOLON,
			wantLit: ";",
		},
		{
			name: "eof",
			fields: fields{
				src:    []byte(`= 1;`),
				offset: 4,
			},
			wantTok: token.EOF,
			wantLit: string(eof),
		},
		{
			name: "str const",
			fields: fields{
				src:    []byte(`= "foo";`),
				offset: 2,
			},
			wantTok: token.STRING_CONST,
			wantLit: "foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				src:    tt.fields.src,
				char:   tt.fields.char,
				offset: tt.fields.offset,
			}
			s.init()
			gotTok, gotLit := s.Scan()
			if !reflect.DeepEqual(gotTok, tt.wantTok) {
				t.Errorf("Scanner.Scan() gotTok = '%v', want '%v'", gotTok, tt.wantTok)
			}
			if gotLit != tt.wantLit {
				t.Errorf("Scanner.Scan() gotLit = '%v', want '%v'", gotLit, tt.wantLit)
			}
		})
	}
}
