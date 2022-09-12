package bomtrim

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

func TestStringAllEncodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{}

	for i := ef; i < endList; i++ {
		tests = append(tests, struct {
			name string
			args args
			want string
		}{name: i.string(), args: args{string(i.hex()) + "a"}, want: "a"},
		)
		tests = append(tests, struct {
			name string
			args args
			want string
		}{name: i.string(), args: args{string(i.hex()) + "apple sauce \n"}, want: "apple sauce \n"},
		)

		if i.string() == "0x73" {
			String(string(i.hex()) + "apple sauce \n")
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.s); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
