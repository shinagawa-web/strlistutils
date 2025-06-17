package strlistutils

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	input := []string{"a", "b", "a", "c", "b"}
	want := []string{"a", "b", "c"}
	got := RemoveDuplicates(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
