package bomtrim

import (
	"unicode/utf8"
)

// FirstIndex removes the BOM from the first index of the given slice.
func FirstIndex(record []string) []string {
	if len(record) == 0 {
		return record
	}

	record[0] = String(record[0])

	return record
}

// String removes the BOM from the given string.
func String(s string) string {
	if len(s) == 0 {
		return s
	}

	utfRune, _ := utf8.DecodeRuneInString(s)

	var new []rune
	removedRune := false

	for _, r := range s {
		if r != utfRune || removedRune {
			new = append(new, r)
		} else {
			removedRune = true
		}
	}

	s = string(new)

	return s
}
