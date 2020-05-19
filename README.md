# splitter

A small Go utility for splitting a file into smaller ones.

## Installation
```shell script
$ go get -u github.com/knwoop/splitter
```

## sample cmd
```shell script
$ splitter -path _example/example.csv -sep 3 -head 
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
    readers, err := splitter.Split(f, true, 40)
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
	readers, err := splitter.Split(r, true, 3)
	if err != nil {
		log.Fatal(err)
	}
	buf1 := new(bytes.Buffer)
	if _, err := buf1.ReadFrom(readers[0]); err != nil {
		log.Fatal(err)
	}
	buf2 := new(bytes.Buffer)
	if _, err := buf2.ReadFrom(readers[1]); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf1)
	fmt.Println(buf2)
}
```
