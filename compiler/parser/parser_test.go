package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/schattian/nand2tetris/compiler/parse"
	"github.com/schattian/nand2tetris/compiler/token"
)

func nodeFromSchema(t *testing.T, schema *nodeSchema, state state) *node {
	t.Helper()
	n := schema.newNode()
	n.state = state
	return n
}

func TestParse(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name   string
		src    []byte
		parent *node
		want   parse.Node
	}{
		{
			name:   "subroutine-body-return-expression",
			src:    []byte(`return (1 + 1);`),
			parent: nodeFromSchema(t, nodeSubroutineBody, 0),
			want: &node{
				closed: true,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.RETURN, Literal: "return"}),
					&node{ // expr
						closed: false,
						children: []parse.Node{
							&node{ // term
								closed: true,
								children: []parse.Node{
									newTokenNode(&parse.Token{Token: token.LPAREN, Literal: "("}),
									&node{ // expr
										closed: false,
										children: []parse.Node{
											&node{ // term
												closed: true,
												children: []parse.Node{
													newTokenNode(&parse.Token{Token: token.INTEGER_CONST, Literal: "1"}),
												},
											},
											newTokenNode(&parse.Token{Token: token.AND, Literal: "+"}),
											&node{ // term
												closed: true,
												children: []parse.Node{
													newTokenNode(&parse.Token{Token: token.INTEGER_CONST, Literal: "1"}),
												},
											},
										},
									},
									newTokenNode(&parse.Token{Token: token.RPAREN, Literal: ")"}),
								},
							},
						},
					},
					newTokenNode(&parse.Token{Token: token.SEMICOLON, Literal: ";"}),
				},
			},
		},
		{
			name:   "subroutine-body-return-expression-indexing",
			src:    []byte(`return foo[1];`),
			parent: nodeFromSchema(t, nodeSubroutineBody, 0),
			want: &node{
				closed: true,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.RETURN, Literal: "return"}),
					&node{ // expr
						closed: false,
						children: []parse.Node{
							&node{ // term
								closed: true,
								children: []parse.Node{
									newTokenNode(&parse.Token{Token: token.IDENT, Literal: "foo"}),
									newTokenNode(&parse.Token{Token: token.LBRACK, Literal: "["}),
									&node{ // expr
										children: []parse.Node{
											&node{ // term
												closed: true,
												children: []parse.Node{
													newTokenNode(&parse.Token{Token: token.INTEGER_CONST, Literal: "1"}),
												},
											},
										},
									},
									newTokenNode(&parse.Token{Token: token.RBRACK, Literal: "]"}),
								},
							},
						},
					},
					newTokenNode(&parse.Token{Token: token.SEMICOLON, Literal: ";"}),
				},
			},
		},
		{
			name:   "subroutine-body-return-expression-subroutineCall",
			src:    []byte(`return foo();`),
			parent: nodeFromSchema(t, nodeSubroutineBody, 0),
			want: &node{
				closed: true,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.RETURN, Literal: "return"}),
					&node{ // expr
						children: []parse.Node{
							&node{ // term
								closed: true,
								children: []parse.Node{
									&node{ // subroutineCall
										closed: true,
										children: []parse.Node{
											newTokenNode(&parse.Token{Token: token.IDENT, Literal: "foo"}),
											newTokenNode(&parse.Token{Token: token.LPAREN, Literal: "("}),
											newTokenNode(&parse.Token{Token: token.RPAREN, Literal: ")"}),
										},
									},
								},
							},
						},
					},
					newTokenNode(&parse.Token{Token: token.SEMICOLON, Literal: ";"}),
				},
			},
		},

		{
			name:   "subroutine-body-return-expression-multiple-terms",
			src:    []byte(`return true & 1 - "baz";`),
			parent: nodeFromSchema(t, nodeSubroutineBody, 0),
			want: &node{
				closed: true,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.RETURN, Literal: "return"}),
					&node{ // expr
						closed: false,
						children: []parse.Node{
							&node{ // term
								closed: true,
								children: []parse.Node{
									newTokenNode(&parse.Token{Token: token.TRUE, Literal: "true"}),
								},
							},
							newTokenNode(&parse.Token{Token: token.AND, Literal: "&"}),

							&node{ // term
								closed: true,
								children: []parse.Node{
									newTokenNode(&parse.Token{Token: token.INTEGER_CONST, Literal: "1"}),
								},
							},
							newTokenNode(&parse.Token{Token: token.SUB, Literal: "-"}),
							&node{ // term
								closed: true,
								children: []parse.Node{
									newTokenNode(&parse.Token{Token: token.STRING_CONST, Literal: "baz"}),
								},
							},
						},
					},
					newTokenNode(&parse.Token{Token: token.SEMICOLON, Literal: ";"}),
				},
			},
		},

		{
			name:   "subroutine-body-return-expression",
			src:    []byte(`return true;`),
			parent: nodeFromSchema(t, nodeSubroutineBody, 0),
			want: &node{
				closed: true,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.RETURN, Literal: "return"}),
					&node{ // expr
						closed: false,
						children: []parse.Node{
							&node{ // term
								closed: true,
								children: []parse.Node{
									newTokenNode(&parse.Token{Token: token.TRUE, Literal: "true"}),
								},
							},
						},
					},
					newTokenNode(&parse.Token{Token: token.SEMICOLON, Literal: ";"}),
				},
			},
		},
		{
			name:   "subroutine-body-return",
			src:    []byte(`return;`),
			parent: nodeFromSchema(t, nodeSubroutineBody, 0),
			want: &node{
				closed: true,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.RETURN, Literal: "return"}),
					newTokenNode(&parse.Token{Token: token.SEMICOLON, Literal: ";"}),
				},
			},
		},

		{
			name:   "subroutine-body-do-func",
			src:    []byte(`do something();`),
			parent: nodeFromSchema(t, nodeSubroutineBody, 0),
			want: &node{
				closed: true,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.DO, Literal: "do"}),
					&node{
						closed: true,
						children: []parse.Node{
							newTokenNode(&parse.Token{Token: token.IDENT, Literal: "something"}),
							newTokenNode(&parse.Token{Token: token.LPAREN, Literal: "("}),
							newTokenNode(&parse.Token{Token: token.RPAREN, Literal: ")"}),
						},
					},
					newTokenNode(&parse.Token{Token: token.SEMICOLON, Literal: ";"}),
				},
			},
		},

		{
			name:   "subroutine-body-do-method",
			src:    []byte(`do fooBar.something();`),
			parent: nodeFromSchema(t, nodeSubroutineBody, 0),
			want: &node{
				closed: true,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.DO, Literal: "do"}),
					&node{
						closed: true,
						children: []parse.Node{
							newTokenNode(&parse.Token{Token: token.IDENT, Literal: "fooBar"}),
							newTokenNode(&parse.Token{Token: token.DOT, Literal: "."}),
							newTokenNode(&parse.Token{Token: token.IDENT, Literal: "something"}),
							newTokenNode(&parse.Token{Token: token.LPAREN, Literal: "("}),
							newTokenNode(&parse.Token{Token: token.RPAREN, Literal: ")"}),
						},
					},
					newTokenNode(&parse.Token{Token: token.SEMICOLON, Literal: ";"}),
				},
			},
		},

		{
			name:   "subroutine-body-let-indexing",
			src:    []byte(`let fooBar[1] = true;`),
			parent: nodeFromSchema(t, nodeSubroutineBody, 0),
			want: &node{
				closed: true,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.LET, Literal: "let"}),
					newTokenNode(&parse.Token{Token: token.IDENT, Literal: "fooBar"}),
					newTokenNode(&parse.Token{Token: token.LBRACK, Literal: "["}),
					&node{ // expr
						closed: false,
						children: []parse.Node{
							&node{ // term
								closed: true,
								children: []parse.Node{
									newTokenNode(&parse.Token{Token: token.INTEGER_CONST, Literal: "1"}),
								},
							},
						},
					},
					newTokenNode(&parse.Token{Token: token.RBRACK, Literal: "]"}),
					newTokenNode(&parse.Token{Token: token.EQ, Literal: "="}),
					&node{ // expr
						closed: false,
						children: []parse.Node{
							&node{ // term
								closed: true,
								children: []parse.Node{
									newTokenNode(&parse.Token{Token: token.TRUE, Literal: "true"}),
								},
							},
						},
					},
					newTokenNode(&parse.Token{Token: token.SEMICOLON, Literal: ";"}),
				},
			},
		},

		{
			name:   "subroutine-body-let",
			src:    []byte(`let fooBar = true;`),
			parent: nodeFromSchema(t, nodeSubroutineBody, 0),
			want: &node{
				closed: true,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.LET, Literal: "let"}),
					newTokenNode(&parse.Token{Token: token.IDENT, Literal: "fooBar"}),
					newTokenNode(&parse.Token{Token: token.EQ, Literal: "="}),

					&node{ // expr
						closed: false,
						children: []parse.Node{
							&node{ // term
								closed: true,
								children: []parse.Node{
									newTokenNode(&parse.Token{Token: token.TRUE, Literal: "true"}),
								},
							},
						},
					},
					newTokenNode(&parse.Token{Token: token.SEMICOLON, Literal: ";"}),
				},
			},
		},

		{
			name:   "subroutine-body-if-else",
			src:    []byte(`if (true) {} else {}`),
			parent: nodeFromSchema(t, nodeSubroutineBody, 0),
			want: &node{
				closed: true,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.IF, Literal: "if"}),
					newTokenNode(&parse.Token{Token: token.LPAREN, Literal: "("}),
					&node{ // expr
						closed: false,
						children: []parse.Node{
							&node{ // term
								closed: true,
								children: []parse.Node{
									newTokenNode(&parse.Token{Token: token.TRUE, Literal: "true"}),
								},
							},
						},
					},
					newTokenNode(&parse.Token{Token: token.ELSE, Literal: "else"}),
					newTokenNode(&parse.Token{Token: token.LBRACE, Literal: "{"}),
					newTokenNode(&parse.Token{Token: token.RBRACE, Literal: "}"}),

					newTokenNode(&parse.Token{Token: token.RPAREN, Literal: ")"}),
					newTokenNode(&parse.Token{Token: token.LBRACE, Literal: "{"}),
					newTokenNode(&parse.Token{Token: token.RBRACE, Literal: "}"}),
				},
			},
		},

		{
			name:   "subroutine-body-if",
			src:    []byte(`if (true) {}`),
			parent: nodeFromSchema(t, nodeSubroutineBody, 0),
			want: &node{
				closed: false,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.IF, Literal: "if"}),
					newTokenNode(&parse.Token{Token: token.LPAREN, Literal: "("}),
					&node{ // expr
						closed: false,
						children: []parse.Node{
							&node{ // term
								closed: true,
								children: []parse.Node{
									newTokenNode(&parse.Token{Token: token.TRUE, Literal: "true"}),
								},
							},
						},
					},
					newTokenNode(&parse.Token{Token: token.RPAREN, Literal: ")"}),
					newTokenNode(&parse.Token{Token: token.LBRACE, Literal: "{"}),
					newTokenNode(&parse.Token{Token: token.RBRACE, Literal: "}"}),
				},
			},
		},

		{
			name:   "subroutine-body-while",
			src:    []byte(`while (true) {}`),
			parent: nodeFromSchema(t, nodeSubroutineBody, 0),
			want: &node{
				closed: true,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.WHILE, Literal: "while"}),
					newTokenNode(&parse.Token{Token: token.LPAREN, Literal: "("}),
					&node{ // expr
						closed: false,
						children: []parse.Node{
							&node{ // term
								closed: true,
								children: []parse.Node{
									newTokenNode(&parse.Token{Token: token.TRUE, Literal: "true"}),
								},
							},
						},
					},
					newTokenNode(&parse.Token{Token: token.RPAREN, Literal: ")"}),
					newTokenNode(&parse.Token{Token: token.LBRACE, Literal: "{"}),
					newTokenNode(&parse.Token{Token: token.RBRACE, Literal: "}"}),
				},
			},
		},
		{
			name: "subroutine-dec-body-varDec",
			src: []byte(`function void fooBar (int qux) {
							var int foo, bar;
						}`),
			parent: nodeFromSchema(t, nodeClass, 0),
			want: &node{
				closed: true,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.FUNCTION, Literal: "function"}),
					newTokenNode(&parse.Token{Token: token.VOID, Literal: "void"}),
					newTokenNode(&parse.Token{Token: token.IDENT, Literal: "fooBar"}),
					newTokenNode(&parse.Token{Token: token.LPAREN, Literal: "("}),
					&node{
						closed: false,
						children: []parse.Node{
							newTokenNode(&parse.Token{Token: token.INT, Literal: "int"}),
							newTokenNode(&parse.Token{Token: token.IDENT, Literal: "qux"}),
						},
					},
					newTokenNode(&parse.Token{Token: token.RPAREN, Literal: ")"}),
					&node{
						closed: true,
						children: []parse.Node{
							newTokenNode(&parse.Token{Token: token.LBRACE, Literal: "{"}),
							&node{
								closed: true,
								children: []parse.Node{
									newTokenNode(&parse.Token{Token: token.VAR, Literal: "var"}),
									newTokenNode(&parse.Token{Token: token.INT, Literal: "int"}),
									newTokenNode(&parse.Token{Token: token.IDENT, Literal: "foo"}),
									newTokenNode(&parse.Token{Token: token.COMMA, Literal: ","}),
									newTokenNode(&parse.Token{Token: token.IDENT, Literal: "bar"}),
									newTokenNode(&parse.Token{Token: token.SEMICOLON, Literal: ";"}),
								},
							},
							newTokenNode(&parse.Token{Token: token.RBRACE, Literal: "}"}),
						},
					},
				},
			},
		},

		{
			name: "class-classVarDec-subroutineDec-subroutineBody-varDec",
			src: []byte(`class Foo {
				static char quz;
				function void fooBar (int qux) {
				}
				}`),
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
							newTokenNode(&parse.Token{Token: token.CHAR, Literal: "char"}),
							newTokenNode(&parse.Token{Token: token.IDENT, Literal: "quz"}),
							newTokenNode(&parse.Token{Token: token.SEMICOLON, Literal: ";"}),
						},
					},

					&node{
						closed: true,
						children: []parse.Node{
							newTokenNode(&parse.Token{Token: token.FUNCTION, Literal: "function"}),
							newTokenNode(&parse.Token{Token: token.VOID, Literal: "void"}),
							newTokenNode(&parse.Token{Token: token.IDENT, Literal: "fooBar"}),
							newTokenNode(&parse.Token{Token: token.LPAREN, Literal: "("}),
							&node{
								closed: false,
								children: []parse.Node{
									newTokenNode(&parse.Token{Token: token.INT, Literal: "int"}),
									newTokenNode(&parse.Token{Token: token.IDENT, Literal: "qux"}),
								},
							},
							newTokenNode(&parse.Token{Token: token.RPAREN, Literal: ")"}),
							&node{
								closed: true,
								children: []parse.Node{
									newTokenNode(&parse.Token{Token: token.LBRACE, Literal: "{"}),

									newTokenNode(&parse.Token{Token: token.RBRACE, Literal: "}"}),
								},
							},
						},
					},
					newTokenNode(&parse.Token{Token: token.RBRACE, Literal: "}"}),
				},
			},
		},
		{
			name: "class-with-subroutineDec-full",
			src: []byte(`class Foo {
				function void fooBar (int qux) {}
				}`),
			want: &node{
				closed: true,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.CLASS, Literal: "class"}),
					newTokenNode(&parse.Token{Token: token.IDENT, Literal: "Foo"}),
					newTokenNode(&parse.Token{Token: token.LBRACE, Literal: "{"}),
					&node{
						closed: true,
						children: []parse.Node{
							newTokenNode(&parse.Token{Token: token.FUNCTION, Literal: "function"}),
							newTokenNode(&parse.Token{Token: token.VOID, Literal: "void"}),
							newTokenNode(&parse.Token{Token: token.IDENT, Literal: "fooBar"}),
							newTokenNode(&parse.Token{Token: token.LPAREN, Literal: "("}),
							&node{
								closed: false,
								children: []parse.Node{
									newTokenNode(&parse.Token{Token: token.INT, Literal: "int"}),
									newTokenNode(&parse.Token{Token: token.IDENT, Literal: "qux"}),
								},
							},
							newTokenNode(&parse.Token{Token: token.RPAREN, Literal: ")"}),
							&node{
								closed: true,
								children: []parse.Node{
									newTokenNode(&parse.Token{Token: token.LBRACE, Literal: "{"}),
									newTokenNode(&parse.Token{Token: token.RBRACE, Literal: "}"}),
								},
							},
						},
					},
					newTokenNode(&parse.Token{Token: token.RBRACE, Literal: "}"}),
				},
			},
		},

		{
			name: "class-with-subroutineDec",
			src: []byte(`class Foo {
				function void fooBar () {}
				}`),
			want: &node{
				closed: true,
				children: []parse.Node{
					newTokenNode(&parse.Token{Token: token.CLASS, Literal: "class"}),
					newTokenNode(&parse.Token{Token: token.IDENT, Literal: "Foo"}),
					newTokenNode(&parse.Token{Token: token.LBRACE, Literal: "{"}),
					&node{
						closed: true,
						children: []parse.Node{
							newTokenNode(&parse.Token{Token: token.FUNCTION, Literal: "function"}),
							newTokenNode(&parse.Token{Token: token.VOID, Literal: "void"}),
							newTokenNode(&parse.Token{Token: token.IDENT, Literal: "fooBar"}),
							newTokenNode(&parse.Token{Token: token.LPAREN, Literal: "("}),
							newTokenNode(&parse.Token{Token: token.RPAREN, Literal: ")"}),
							&node{
								closed: true,
								children: []parse.Node{
									newTokenNode(&parse.Token{Token: token.LBRACE, Literal: "{"}),
									newTokenNode(&parse.Token{Token: token.RBRACE, Literal: "}"}),
								},
							},
						},
					},
					newTokenNode(&parse.Token{Token: token.RBRACE, Literal: "}"}),
				},
			},
		},
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
			got := New(tt.src, tt.parent).Parse()
			diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(node{}, nodeSchema{}), cmpopts.IgnoreFields(node{}, "schema", "state", "fieldsBySubset", "lastFieldSubset"))
			if diff != "" {
				t.Errorf("mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
