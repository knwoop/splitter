package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var path, out string
	var sep int
	var head bool

	flag.StringVar(&path, "path", "", "set csv file path")
	flag.StringVar(&out, "dir", "", "out to dir")
	flag.IntVar(&sep, "sep", 0, "number to split a file")
	flag.BoolVar(&head, "head", false, "have a header in csv ?")
	flag.Parse()

	f, err := os.Open(path)
	if err != nil {
		log.Fatal("error open a file", err.Error())
	}
	readers, err := Split(f, head, sep)
	if err != nil {
		log.Fatal(err)
	}

	for i, r := range readers {
		fileName := filepath.Base(path)
		if err := writeFile(fmt.Sprintf("%d-%s", i, fileName), r); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("success to split a file")
}

func writeFile(fileName string, r io.Reader) error {
	fp, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer fp.Close()

	if _, err := io.Copy(fp, r); err != nil {
		return err
	}
	return nil
}
