// Package strlistutils provides utility functions for working with string slices.
package strlistutils

import "strings"

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

// TrimEach returns a new slice where each string is trimmed of leading and trailing whitespace.
// The original slice is not modified.
func TrimEach(input []string) []string {
	result := make([]string, len(input))
	for i, s := range input {
		result[i] = strings.TrimSpace(s)
	}
	return result
}

// FilterEmpty returns a new slice with all empty strings removed from the input.
// The original slice is not modified.
func FilterEmpty(input []string) []string {
	result := make([]string, 0, len(input))
	for _, s := range input {
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}
