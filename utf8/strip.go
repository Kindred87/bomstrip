package utf8

import "unicode/utf8"

// StripFirstIndex removes UTF8 encoding from the first index of the given slice.
func StripFirstIndex(record []string) []string {
	record[0] = StripLine(record[0])

	return record
}

// StripLine removes UTF8 encoding from the given string.
func StripLine(s string) string {
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
