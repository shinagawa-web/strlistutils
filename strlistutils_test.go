package strlistutils

import (
	"reflect"
	"strings"
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

func TestMap(t *testing.T) {
	input := []string{"  apple ", " banana", "orange "}
	expected := []string{"apple", "banana", "orange"}

	result := Map(input, strings.TrimSpace)

	if len(result) != len(expected) {
		t.Fatalf("Map result length = %d; want %d", len(result), len(expected))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Map result[%d] = %q; want %q", i, result[i], expected[i])
		}
	}
}

func TestFilter(t *testing.T) {
	input := []string{"apple", "banana", "", "orange", ""}
	expected := []string{"apple", "banana", "orange"}
	result := Filter(input, func(s string) bool { return s != "" })
	if len(result) != len(expected) {
		t.Fatalf("Filter length = %d; want %d", len(result), len(expected))
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Filter result[%d] = %q; want %q", i, result[i], expected[i])
		}
	}
}

func TestJoin(t *testing.T) {
	input := []string{"apple", "banana", "orange"}
	expected := "apple,banana,orange"
	result := Join(input, ",")
	if result != expected {
		t.Errorf("Join(%v, %q) = %q; want %q", input, ",", result, expected)
	}
}

func TestReverse(t *testing.T) {
	input := []string{"apple", "banana", "orange"}
	expected := []string{"orange", "banana", "apple"}
	result := Reverse(input)
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Reverse result[%d] = %q; want %q", i, result[i], expected[i])
		}
	}
}
