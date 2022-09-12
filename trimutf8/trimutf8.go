package trimutf8

import (
	"unicode/utf8"
)

// FirstIndex removes UTF8 encoding from the first index of the given slice.
func FirstIndex(record []string) []string {
	record[0] = Line(record[0])

	return record
}

// Line removes UTF8 encoding from the given string.
func Line(s string) string {
	if len(s) == 0 {
		return s
	}

	_, i := utf8.DecodeRuneInString(s)

	t := s[i+1:]

	_, i = utf8.DecodeRuneInString(t)
	t = t[:len(t)-i]

	s = t

	return s
}
