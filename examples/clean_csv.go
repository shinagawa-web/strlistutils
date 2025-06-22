package main

import (
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/shinagawa-web/strlistutils"
)

func main() {
	rawCSV := ` apple , banana , , orange 
grape , , melon , apple `

	r := csv.NewReader(strings.NewReader(rawCSV))
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	for i, row := range records {
		fmt.Printf("Original Row %d: %#v\n", i+1, row)

		// Trim each value
		trimmed := strlistutils.TrimEach(row)

		// Remove empty entries
		filtered := strlistutils.FilterEmpty(trimmed)

		// Remove duplicates
		unique := strlistutils.RemoveDuplicates(filtered)

		fmt.Printf("Cleaned Row %d: %#v\n", i+1, unique)
	}
}
