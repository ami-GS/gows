package gows

import (
	"reflect"
	"testing"
)

func TestNewFrame(t *testing.T) {
	actual := NewFrame()
	expect := &Frame{0, [3]byte{0, 0, 0}, CONTINUE, 0, 0, [4]byte{0, 0, 0, 0}, []byte{}}

	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("got %v\nwant %v", actual, expect)
	}
}
