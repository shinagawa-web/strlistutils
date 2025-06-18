package strlistutils

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func BenchmarkRemoveDuplicates(b *testing.B) {
	input := make([]string, 1000)
	for i := range input {
		input[i] = fmt.Sprintf("item%d", i%100) // duplicates every 100 items
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = RemoveDuplicates(input)
	}
}

func BenchmarkTrimEach(b *testing.B) {
	input := make([]string, 1000)
	for i := range input {
		input[i] = "   word   "
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = TrimEach(input)
	}
}

func BenchmarkFilterEmpty(b *testing.B) {
	input := make([]string, 1000)
	for i := range input {
		if i%10 == 0 {
			input[i] = ""
		} else {
			input[i] = "value"
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FilterEmpty(input)
	}
}

func BenchmarkMap(b *testing.B) {
	list := make([]string, 1000)
	for i := range list {
		list[i] = "text"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Map(list, strings.ToUpper)
	}
}

func BenchmarkFilter(b *testing.B) {
	input := make([]string, 1000)
	for i := range input {
		if i%10 == 0 {
			input[i] = ""
		} else {
			input[i] = "data"
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Filter(input, func(s string) bool { return s != "" })
	}
}

func BenchmarkJoin(b *testing.B) {
	input := make([]string, 1000)
	for i := range input {
		input[i] = "word"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(input, ",")
	}
}

func BenchmarkReverse(b *testing.B) {
	input := make([]string, 1000)
	for i := range input {
		input[i] = fmt.Sprintf("%d", rand.Int())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Reverse(input)
	}
}
