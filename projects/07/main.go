package main

import (
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("filename not given")
	}
	filename := os.Args[1]
	r, err := os.Open(filename)
	if err != nil {
		log.Fatalf("os.Open: %v", err)
	}
	defer r.Close()
	moduleName := strings.TrimSuffix(filename, ".vm")
	w, err := os.Create(moduleName + ".asm")
	if err != nil {
		log.Fatalf("os.Create: %v", err)
	}
	defer w.Close()

	d := NewVMDecoder(moduleName, r)
	e := NewASMEncoder(w)
	for {
		vmc, err := d.Decode()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatalf("e.Decode: %v", err)
		}
		if vmc == nil {
			continue
		}
		err = e.Encode(vmc)
		if err != nil {
			log.Fatalf("e.Encode: %v", err)
		}
	}
}
