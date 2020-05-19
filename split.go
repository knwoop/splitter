package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func Split(r io.Reader, hasHeader bool, sep int) ([]io.Reader, error) {
	scanner := bufio.NewScanner(r)
	var b1, b2 []byte
	if hasHeader {
		if !scanner.Scan() {
			return nil, fmt.Errorf("error no data")
		}
		l := scanner.Bytes()
		b1 = append(append(b1, l...), '\n')
		b2 = append(append(b2, l...), '\n')
	}

	dst1 := 0
	dst2 := 0
	for {
		if !scanner.Scan() {
			break
		}
		l := scanner.Bytes()
		if dst1 < sep {
			b1 = append(append(b1, l...), '\n')
			dst1++
		} else {
			b2 = append(append(b2, l...), '\n')
			dst2++
		}
	}

	return []io.Reader{bytes.NewReader(b1), bytes.NewReader(b2)}, nil
}
