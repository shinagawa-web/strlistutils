package strlistutils_test

import (
	"fmt"
	"strings"

	"github.com/shinagawa-web/strlistutils"
)

func ExampleRemoveDuplicates() {
	input := []string{"apple", "banana", "apple", "orange", "banana"}
	output := strlistutils.RemoveDuplicates(input)
	fmt.Println(output)

	// Output: [apple banana orange]
}

func ExampleTrimEach() {
	input := []string{"  apple ", " banana", "orange "}
	output := strlistutils.TrimEach(input)
	fmt.Println(output)

	// Output: [apple banana orange]
}

func ExampleFilterEmpty() {
	input := []string{"apple", "", "banana", "", "orange"}
	output := strlistutils.FilterEmpty(input)
	fmt.Println(output)

	// Output: [apple banana orange]
}

func ExampleMap() {
	input := []string{"  apple ", " banana", "orange "}
	output := strlistutils.Map(input, strings.TrimSpace)
	fmt.Println(output)

	// Output:
	// [apple banana orange]
}

func ExampleFilter() {
	input := []string{"apple", "banana", "", "orange", ""}
	output := strlistutils.Filter(input, func(s string) bool { return s != "" })
	fmt.Println(output)
	// Output: [apple banana orange]
}

func ExampleJoin() {
	input := []string{"apple", "banana", "orange"}
	output := strlistutils.Join(input, ",")
	fmt.Println(output)
	// Output: apple,banana,orange
}

func ExampleReverse() {
	input := []string{"apple", "banana", "orange"}
	output := strlistutils.Reverse(input)
	fmt.Println(output)
	// Output: [orange banana apple]
}
