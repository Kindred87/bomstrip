package bomtrim

import (
	"reflect"
	"testing"
)

func Test_hexes_hex(t *testing.T) {
	for i := ef; i < endList; i++ {
		if reflect.DeepEqual(i.hex(), []byte("")) {
			t.Errorf("hex at index %d is unsupported by hex func", i)
		}
	}
}

func Test_hexes_string(t *testing.T) {
	for i := ef; i < endList; i++ {
		if i.string() == "" {
			t.Errorf("hex at index %d is unsupported by string func", i)
		}
	}
}
