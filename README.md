# strlistutils

![Test](https://github.com/shinagawa-web/strlistutils/actions/workflows/test.yml/badge.svg)
[![codecov](https://codecov.io/gh/shinagawa-web/strlistutils/graph/badge.svg?token=D76743O8J1)](https://codecov.io/gh/shinagawa-web/strlistutils)
[![Go Report Card](https://goreportcard.com/badge/github.com/shinagawa-web/strlistutils)](https://goreportcard.com/report/github.com/shinagawa-web/strlistutils)
[![Go Reference](https://pkg.go.dev/badge/github.com/shinagawa-web/strlistutils.svg)](https://pkg.go.dev/github.com/shinagawa-web/strlistutils)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)


🧰 Utility functions for working with []string in Go.  
📦 Documentation: https://pkg.go.dev/github.com/shinagawa-web/strlistutils

## Features

- 🧹 `RemoveDuplicates` — remove duplicate strings
- ✂️ `TrimEach` — trim whitespace from each element
- ❌ `FilterEmpty` — remove empty strings
- 🔁 `Map` — apply a function to each element
- 🧪 `Filter` — filter elements by a predicate
- 🔗 `Join` — join elements with a separator
- ↩️ `Reverse` — reverse the list

## Installation

```bash
go get github.com/shinagawa-web/strlistutils@v0.3.0
```

🔒 It is recommended to pin a version to avoid unexpected breaking changes.

## Usage

```go
package main

import (
    "fmt"
    "github.com/shinagawa-web/strlistutils"
)

func main() {
    input := []string{" a ", "b", "a", " ", "c"}
    result := strlistutils.RemoveDuplicates(input)
    fmt.Println(result) // Output: [" a ", "b", " ", "c"]
}
```

➡️ For more examples, see [GoDoc](https://pkg.go.dev/github.com/shinagawa-web/strlistutils) or [examples](https://github.com/shinagawa-web/strlistutils/blob/main/strlistutils_example_test.go)

## Available Functions

| Name               | Description                                      |
| ------------------ | ------------------------------------------------ |
| `RemoveDuplicates` | Removes duplicate elements, preserving order     |
| `TrimEach`         | Trims leading/trailing whitespace from each item |
| `FilterEmpty`      | Removes empty strings                            |
| `Map`              | Applies a function to each element               |
| `Filter`           | Keeps only elements matching a predicate         |
| `Join`             | Concatenates list elements with a separator      |
| `Reverse`          | Reverses the order of elements in the list       |

## What's Included

### Test Coverage

- Unit tests for all functions
- Benchmark tests
- Fuzz tests
- GitHub Actions CI
- Codecov integration

## Usage Examples

- Example code in examples/ directory
- Real-world use case: cleaning CSV input


## License

MIT
