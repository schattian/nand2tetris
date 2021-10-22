package parser

import (
	"github.com/schattian/nand2tetris/compiler/parse"
	"github.com/schattian/nand2tetris/compiler/scanner"
	"github.com/schattian/nand2tetris/compiler/token"
)

type parser struct {
	m map[token.Token]parserFunc

	s         *scanner.Scanner
	token     *parse.Token
	prevToken *parse.Token
	nextToken *parse.Token
	parent    *node
}

func New(src []byte) *parser {
	s := scanner.New(src)
	schema := &nodeSchema{}
	return &parser{s: s, parent: schema.newNode()}
}

func new(src []byte, parent *node) *parser {
	s := scanner.New(src)
	if parent == nil {
		schema := &nodeSchema{}
		parent = schema.newNode()
	}
	return &parser{s: s, parent: parent}
}

type parserFunc func() parse.Node

func (p *parser) ParseTree() *parse.Tree {
	root := p.Parse()
	return &parse.Tree{Name: "tree", Root: root}
}

func (p *parser) Parse() parse.Node {
	p.next()
	if p.token.Token == token.EOF {
		return nil
	}

	schema := p.parent.NextNodeSchema(p.token)
	if schema == nil {
		return p.node()
	}
	return p.parseSchema(schema)
}

func (p *parser) prepTermChild(n *node) *node {
	n.AddNode(p.parent.popChild())
	n.AddNode(p.node())
	return n
}

func (p *parser) prepExpression(n *node) *node {
	childSchema := n.NextNodeSchema(p.token)
	n.AddNode(p.parseSchema(childSchema))
	return n
}

func (p *parser) prepBaseCase(n *node) *node {
	n.AddNode(p.node())
	return n
}

func (p *parser) needLookahead(schema *nodeSchema) bool {
	return p.parent.Schema.NodeType == parse.NodeTerm && schema.NodeType == parse.NodeSubroutineCall
}

func (p *parser) parseSchema(schema *nodeSchema) parse.Node {
	n := schema.newNode()
	if p.needLookahead(schema) {
		n = p.prepTermChild(n)
	} else if schema.NodeType == parse.NodeExpression || schema.NodeType == parse.NodeExpressionList {
		n = p.prepExpression(n)
	} else {
		n = p.prepBaseCase(n)
	}
	for {
		p.parent = n
		v := p.Parse()
		ok := n.AddNode(v)
		if !ok {
			p.prev()
		}
		if !ok || n.Closed {
			return n
		}
	}
}

func (p *parser) node() *node {
	return newTokenNode(p.token)
}

func (p *parser) next() {
	p.prevToken = p.token
	if p.nextToken != nil {
		p.token = p.nextToken
		p.nextToken = nil
		return
	}
	p.token = parse.NewToken(p.s.Scan())
}

func (p *parser) prev() {
	if p.prevToken == nil {
		return
	}
	p.nextToken = p.token
	p.token = p.prevToken
	p.prevToken = nil
}
