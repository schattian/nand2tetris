package parser

import (
	"github.com/schattian/nand2tetris/compiler/parse"
	"github.com/schattian/nand2tetris/compiler/scanner"
	"github.com/schattian/nand2tetris/compiler/token"
)

type parser struct {
	m map[token.Token]parserFunc

	s     *scanner.Scanner
	token *parse.Token
}

func Parse(src []byte) parse.Node {
	s := scanner.New(src)
	p := &parser{s: s}
	return p.parse()
}

type parserFunc func() parse.Node

func (p *parser) parse() parse.Node {
	p.next()
	if p.token.Token == token.EOF {
		return nil
	}

	parserFunc, ok := map[token.Token]parserFunc{
		token.CLASS:  p.parseClass,
		token.STATIC: p.parseStatic,
	}[p.token.Token]
	if ok {
		return parserFunc()
	}

	return p.node()
}

func (p *parser) parseStatic() parse.Node {
	n := nodeClassVarDec.newNode()
	n.AddNode(p.node())
	for {
		ok := n.AddNode(p.parse())
		if !ok || n.closed {
			return n
		}
	}
}

func (p *parser) parseClass() parse.Node {
	n := nodeClass.newNode()
	n.AddNode(p.node())
	for {
		ok := n.AddNode(p.parse())
		if !ok {
			return n
		}
	}
}
func (p *parser) node() *node {
	return newTokenNode(p.token)
}

func (p *parser) next() {
	p.token = parse.NewToken(p.s.Scan())
}
