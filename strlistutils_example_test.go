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
