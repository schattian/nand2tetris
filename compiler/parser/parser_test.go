package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/schattian/nand2tetris/compiler/parse"
	"github.com/schattian/nand2tetris/compiler/token"
)

func TestParse(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		src  []byte
		want parse.Node
	}{
		{
			name: "class-with-multiple-classVarDec",
			src:  []byte(`class Foo { static int foo; static int bar, baz; }`),
			want: &node{
				closed: true,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.CLASS, Literal: "class"}),
					newTokenNode(&parse.Token{Token: token.IDENT, Literal: "Foo"}),
					newTokenNode(&parse.Token{Token: token.LBRACE, Literal: "{"}),
					&node{
						closed: true,
						children: []parse.Node{
							newTokenNode(&parse.Token{Token: token.STATIC, Literal: "static"}),
							newTokenNode(&parse.Token{Token: token.INT, Literal: "int"}),
							newTokenNode(&parse.Token{Token: token.IDENT, Literal: "foo"}),
							newTokenNode(&parse.Token{Token: token.SEMICOLON, Literal: ";"}),
						},
					},
					&node{
						closed: true,
						children: []parse.Node{
							newTokenNode(&parse.Token{Token: token.STATIC, Literal: "static"}),
							newTokenNode(&parse.Token{Token: token.INT, Literal: "int"}),
							newTokenNode(&parse.Token{Token: token.IDENT, Literal: "bar"}),
							newTokenNode(&parse.Token{Token: token.COMMA, Literal: ","}),
							newTokenNode(&parse.Token{Token: token.IDENT, Literal: "baz"}),
							newTokenNode(&parse.Token{Token: token.SEMICOLON, Literal: ";"}),
						},
					},
					newTokenNode(&parse.Token{Token: token.RBRACE, Literal: "}"}),
				},
			},
		},
		{
			name: "class-with-classVarDec",
			src:  []byte(`class Foo { static int foo; }`),
			want: &node{
				closed: true,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.CLASS, Literal: "class"}),
					newTokenNode(&parse.Token{Token: token.IDENT, Literal: "Foo"}),
					newTokenNode(&parse.Token{Token: token.LBRACE, Literal: "{"}),
					&node{
						closed: true,
						children: []parse.Node{
							newTokenNode(&parse.Token{Token: token.STATIC, Literal: "static"}),
							newTokenNode(&parse.Token{Token: token.INT, Literal: "int"}),
							newTokenNode(&parse.Token{Token: token.IDENT, Literal: "foo"}),
							newTokenNode(&parse.Token{Token: token.SEMICOLON, Literal: ";"}),
						},
					},
					newTokenNode(&parse.Token{Token: token.RBRACE, Literal: "}"}),
				},
			},
		},
		{
			name: "classVarDec-full",
			src:  []byte(`static int foo, bar;`),
			want: &node{
				closed: true,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.STATIC, Literal: "static"}),
					newTokenNode(&parse.Token{Token: token.INT, Literal: "int"}),
					newTokenNode(&parse.Token{Token: token.IDENT, Literal: "foo"}),
					newTokenNode(&parse.Token{Token: token.COMMA, Literal: ","}),
					newTokenNode(&parse.Token{Token: token.IDENT, Literal: "bar"}),
					newTokenNode(&parse.Token{Token: token.SEMICOLON, Literal: ";"}),
				},
			},
		},
		{
			name: "classVarDec-shallow",
			src:  []byte(`static int foo;`),
			want: &node{
				closed: true,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.STATIC, Literal: "static"}),
					newTokenNode(&parse.Token{Token: token.INT, Literal: "int"}),
					newTokenNode(&parse.Token{Token: token.IDENT, Literal: "foo"}),
					newTokenNode(&parse.Token{Token: token.SEMICOLON, Literal: ";"}),
				},
			},
		},
		{
			name: "class-shallow",
			src:  []byte(`class Foo {}`),
			want: &node{
				closed: true,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.CLASS, Literal: "class"}),
					newTokenNode(&parse.Token{Token: token.IDENT, Literal: "Foo"}),
					newTokenNode(&parse.Token{Token: token.LBRACE, Literal: "{"}),
					newTokenNode(&parse.Token{Token: token.RBRACE, Literal: "}"}),
				},
			},
		},
		{
			name: "no src",
			src:  []byte(""),
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Parse(tt.src)
			diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(node{}, nodeSchema{}), cmpopts.IgnoreFields(node{}, "schema", "fieldsBySubset", "lastFieldSubset"))
			if diff != "" {
				t.Errorf("mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
