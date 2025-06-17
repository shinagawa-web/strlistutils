// Package strlistutils provides utility functions for working with string slices.
package strlistutils

// RemoveDuplicates returns a new slice with duplicate strings removed from the input.
// The order of the first occurrence of each element is preserved.
// This function does not modify the original slice.
func RemoveDuplicates(input []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, s := range input {
		if !seen[s] {
			seen[s] = true
			result = append(result, s)
		}
	}
	return result
}
