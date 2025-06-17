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

func TestTrimEach(t *testing.T) {
	input := []string{"  apple ", "\tbanana\t", "\norange\n"}
	want := []string{"apple", "banana", "orange"}
	got := TrimEach(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFilterEmpty(t *testing.T) {
	input := []string{"apple", "", "banana", "", "orange"}
	want := []string{"apple", "banana", "orange"}
	got := FilterEmpty(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
