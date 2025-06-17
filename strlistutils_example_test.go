package strlistutils_test

import (
	"fmt"

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
