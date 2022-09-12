package trimutf8

import (
	"strings"
	"testing"
)

func TestFirstIndex(t *testing.T) {
	var record []string
	record = append(record, "\uFEFFa")
	record = append(record, "\uFEFFb")

	record = FirstIndex(record)

	if strings.ContainsRune(record[0], []rune("\uFEFF")[0]) {
		t.Fatalf("First index contains BOM after trimming")
	} else if !strings.ContainsRune(record[1], []rune("\uFEFF")[0]) {
		t.Fatalf("Second index was trimmed of BOM when only the first index was supposed to be trimmed")
	}
}

func TestLinePrefix(t *testing.T) {
	s := "\uFEFFa"

	s = String(s)

	if strings.ContainsRune(s, []rune("\uFEFF")[0]) {
		t.Fatalf("String contains BOM after trimming")
	}
}

func TestLinePrefixAndSuffix(t *testing.T) {
	s := "\uFEFFa\uFEFF"

	s = String(s)

	if strings.ContainsRune(s, []rune("\uFEFF")[0]) {
		t.Fatalf("String contains BOM after trimming")
	}
}
