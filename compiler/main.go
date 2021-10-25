package main

import (
	"encoding/xml"
	"github.com/schattian/nand2tetris/compiler/parse/parser"
	"log"
	"os"
)

func main() {
	srcFilename, dstFilename := os.Args[1], os.Args[2]

	src, err := os.ReadFile(srcFilename)
	if err != nil {
		log.Fatal(err)
	}

	w, err := os.Create(dstFilename)
	if err != nil {
		log.Fatal(err)
	}

	tree := parser.New(src).ParseTree()
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	err = enc.Encode(tree)
	if err != nil {
		log.Fatal(err)
	}
}
