package main

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("not enough args")
	}
	destFilename := os.Args[2]

	w, err := os.Create(destFilename)
	if err != nil {
		log.Fatalf("os.Create: %v", err)
	}
	defer w.Close()
	enc, err := NewASMEncoder(w)
	if err != nil {
		log.Fatalf("NewASMEncoder: %v", err)
	}

	if filepath.Ext(os.Args[1]) != ".vm" {
		err = translateDir(os.Args[1], enc)
	} else {
		err = translateFile(os.Args[1], enc)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func translateDir(dirname string, enc *ASMEncoder) error {
	entries, err := os.ReadDir(dirname)
	if err != nil {
		return err
	}
	var filenames []string
	for _, entry := range entries {
		if filepath.Ext(entry.Name()) == ".vm" {
			filenames = append(filenames, entry.Name())
		}
	}
	if err != nil {
		return err
	}
	for _, filename := range filenames {
		err = translateFile(dirname+"/"+filename, enc)
		if err != nil {
			return err
		}
	}
	return nil
}

func translateFile(filename string, enc *ASMEncoder) error {
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer r.Close()

	d := NewVMDecoder(filepath.Base(filename), r)
	for {
		vmc, err := d.Decode()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}
		if vmc == nil {
			continue
		}
		err = enc.Encode(vmc)
		if err != nil {
			return err
		}
	}
	return nil
}
