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

// Contains returns true if the target string is present in the list.
func Contains(list []string, target string) bool {
	for _, s := range list {
		if s == target {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the target string in the list.
// If not found, it returns -1.
func IndexOf(list []string, target string) int {
	for i, s := range list {
		if s == target {
			return i
		}
	}
	return -1
}

// Map applies a function to each string in the list and returns a new list.
func Map(list []string, f func(string) string) []string {
	result := make([]string, len(list))
	for i, s := range list {
		result[i] = f(s)
	}
	return result
}

// Filter returns a new list containing only the elements for which the predicate returns true.
func Filter(list []string, predicate func(string) bool) []string {
	result := make([]string, 0, len(list))
	for _, s := range list {
		if predicate(s) {
			result = append(result, s)
		}
	}
	return result
}

// Join joins all elements of the list into a single string with the given separator.
func Join(list []string, sep string) string {
	return strings.Join(list, sep)
}

// Reverse returns a new list with the elements in reverse order.
func Reverse(list []string) []string {
	n := len(list)
	result := make([]string, n)
	for i, s := range list {
		result[n-1-i] = s
	}
	return result
}
