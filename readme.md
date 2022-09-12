# bomtrim

A tiny utility that removes byte order marks so you don't have to!  

Byte order marks are special unicode characters prefixed to text files, such as \uFEFF.  

<br>

# Install
```
GOPROXY=direct go get -u github.com/Kindred87/bomtrim
```

<br>

## Quickstart
```go
package main

import (
    "encoding/csv"
    "fmt"
    "os"

    "github.com/Kindred87/bomtrim"
)

func main() {
    fi, err := os.Open("foo.csv")
    if err != nil {
        log.Fatalf("error while reading foo.csv: %s", err.Error())
    }

    defer fi.Close()

    reader := csv.NewReader(fi)

    firstRow, err := reader.Read()
	if err != nil {
		log.Fatalf("error while reading header row in %s: %s", fi.Name(), err.Error())
	}

    fmt.Printf("Before trimming: %#v\n", firstRow)

    firstRow = bomtrim.FirstIndex(record []string)

    fmt.Printf("After trimming: %#v\n", firstRow)
}
```