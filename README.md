# splitter

A small Go utility for splitting a file into smaller ones.

## Installation
```shell script
$ go get -u github.com/knwoop/splitter
```

## how to use
```go
package main

import (
    "fmt"
    "log"
    "os"
    
    "github.com/knwoop/splitter"
)

func main() {
    f, err := os.Open("test.csv")
    if err != nil{
        log.Fatal(err)
    }
    defer f.Close() 
    b1, b2, err := splitter.Split(f, true, 40, 60)
    if err != nil {
        log.Fatal(err)
    }
    ...
}
```

## Let's playground
```go
package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/knwoop/splitter"
)

var csv = `header1,header2
"1111","22222"
33333,44444
55555,66666
77777,88888`

func main() {
	r := strings.NewReader(csv)
	b1, b2, err := splitter.Split(r, true, 1, 3)
	if err != nil {
		log.Fatal(err)
	}
	buf1 := new(bytes.Buffer)
	if _, err := buf1.ReadFrom(b1); err != nil {
		log.Fatal(err)
	}
	buf2 := new(bytes.Buffer)
	if _, err := buf2.ReadFrom(b2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf1)
	fmt.Println(buf2)
}
```
