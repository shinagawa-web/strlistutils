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
- 🔍 `Contains` — check if a value exists
- 🔢 `IndexOf` — find the index of a string

## Installation

```bash
go get github.com/shinagawa-web/strlistutils@v0.2.0
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

| Name | Description |
|------|-------------|
| `RemoveDuplicates` | Removes duplicate elements, preserving order |
| `TrimEach` | Trims leading/trailing whitespace from each string |
| `FilterEmpty` | Removes empty strings |
| `Map` | Applies a function to each element |
| `Filter` | Keeps elements that match a condition |

## Roadmap

### ✅ Core Utilities

- [x] RemoveDuplicates
- [x] TrimEach
- [x] FilterEmpty
- [x] Map
- [x] Filter
- [x] Join
- [x] Reverse

### 🧪 Testing & Quality

- [x] GitHub Actions (CI)
- [x] Codecov integration
- [x] Add Fuzz testing
- [x] Add Benchmark tests

### 📚 Documentation

- [x] GoDoc examples
- [x] Function usage table
- [ ] Add real-world usage examples
- [ ] Write Zenn/Qiita article

### 🚀 Release Milestones

- [x] v0.1.0: Initial release with basic utilities
- [x] v0.2.0: Expanded with more functions and CI integration
- [ ] v0.3.0: Add Map / Filter / Join and prepare for feature completeness
- [ ] v1.0.0: API freeze, complete documentation, production-ready release

## License

MIT
