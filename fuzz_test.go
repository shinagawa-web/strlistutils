package strlistutils

import (
	"strings"
	"testing"
)

func FuzzRemoveDuplicates(f *testing.F) {
	f.Add("apple banana apple orange banana")
	f.Fuzz(func(t *testing.T, input string) {
		slice := strings.Fields(input)
		result := RemoveDuplicates(slice)
		seen := make(map[string]bool)
		for _, s := range result {
			if seen[s] {
				t.Errorf("duplicate found in result: %q", s)
			}
			seen[s] = true
		}
	})
}

func FuzzTrimEach(f *testing.F) {
	f.Add("  apple ", " banana", "orange ")
	f.Fuzz(func(t *testing.T, a, b, c string) {
		input := []string{a, b, c}
		result := TrimEach(input)
		if len(result) != len(input) {
			t.Errorf("TrimEach: length mismatch, got %d, want %d", len(result), len(input))
		}
		for _, s := range result {
			if strings.HasPrefix(s, " ") || strings.HasSuffix(s, " ") {
				t.Errorf("TrimEach failed to trim: %q", s)
			}
		}
	})
}

func FuzzMap(f *testing.F) {
	f.Add("apple", "banana", "orange")
	f.Fuzz(func(t *testing.T, a, b, c string) {
		input := []string{a, b, c}
		mapped := Map(input, strings.ToUpper)
		if len(mapped) != len(input) {
			t.Errorf("Map length mismatch: got %d, want %d", len(mapped), len(input))
		}
		for i := range input {
			if mapped[i] != strings.ToUpper(input[i]) {
				t.Errorf("Map mismatch at index %d: got %q, want %q", i, mapped[i], strings.ToUpper(input[i]))
			}
		}
	})
}

func FuzzFilterEmpty(f *testing.F) {
	f.Add("apple", "", "banana")
	f.Fuzz(func(t *testing.T, a, b, c string) {
		input := []string{a, b, c}
		result := FilterEmpty(input)
		for _, s := range result {
			if s == "" {
				t.Error("FilterEmpty did not filter out empty string")
			}
		}
	})
}

func FuzzFilter(f *testing.F) {
	f.Add("apple", "banana", "orange")
	f.Fuzz(func(t *testing.T, a, b, c string) {
		input := []string{a, b, c}

		result := Filter(input, func(s string) bool {
			return len(s) >= 5
		})

		for _, s := range result {
			if len(s) < 5 {
				t.Errorf("Filter failed: %q has length < 5", s)
			}
		}
	})
}

func FuzzJoin(f *testing.F) {
	f.Add("apple banana orange", ",")
	f.Fuzz(func(t *testing.T, input, sep string) {
		slice := strings.Fields(input)
		expected := strings.Join(slice, sep)
		if Join(slice, sep) != expected {
			t.Errorf("Join mismatch: got %q, want %q", Join(slice, sep), expected)
		}
	})
}

func FuzzReverse(f *testing.F) {
	f.Add("apple", "banana", "orange")
	f.Fuzz(func(t *testing.T, a, b, c string) {
		input := []string{a, b, c}
		rev := Reverse(input)
		rev2 := Reverse(rev)
		for i := range input {
			if input[i] != rev2[i] {
				t.Errorf("Reverse is not symmetric: got %v, want %v", rev2, input)
				break
			}
		}
	})
}
