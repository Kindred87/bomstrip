# bomstrip

Utility library that removes byte-order marks so you don't have to!

<br>

# Install
```
GOPROXY=direct go get -u github.com/Kindred87/bomstrip
```

<br>

## Quickstart
```go
package main

import (
    "encoding/csv"
	"fmt"
	"os"

    "github.com/Kindred87/bomstrip/striputf8"
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

    fmt.Printf("Before stripping: %#v\n", firstRow)

    firstRow = striputf8.FirstIndex(record []string)

    fmt.Printf("After stripping: %#v\n", firstRow)
}
```