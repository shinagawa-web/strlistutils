# strlistutils

![Test](https://github.com/shinagawa-web/strlistutils/actions/workflows/test.yml/badge.svg)
[![codecov](https://codecov.io/gh/shinagawa-web/strlistutils/graph/badge.svg?token=D76743O8J1)](https://codecov.io/gh/shinagawa-web/strlistutils)
[![Go Report Card](https://goreportcard.com/badge/github.com/shinagawa-web/strlistutils)](https://goreportcard.com/report/github.com/shinagawa-web/strlistutils)
[![Go Reference](https://pkg.go.dev/badge/github.com/shinagawa-web/strlistutils.svg)](https://pkg.go.dev/github.com/shinagawa-web/strlistutils)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)


🧰 Utility functions for working with []string in Go.  
📦 Documentation: https://pkg.go.dev/github.com/shinagawa-web/strlistutils


## Why

Developers repeatedly re-implement the same []string routines (dedup, trim, filter).
This package centralizes those patterns with clear behavior and tests, so teams can:

- avoid ad-hoc snippets scattered across services,
- keep behavior consistent and reviewable,
- get predictable performance without external deps.

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

## Install

```sh
go get github.com/shinagawa-web/strlistutils@v0.3.0
```

Pin a version to avoid accidental breakage.

## Quick Start

```go
package main

import (
	"fmt"

	"github.com/shinagawa-web/strlistutils"
)

func main() {
	input := []string{" a ", "b", "a", " ", "c"}

	out := strlistutils.RemoveDuplicates(
		strlistutils.FilterEmpty(
			strlistutils.TrimEach(input),
		),
	)
	fmt.Println(out) // ["a", "b", "c"]
}
```

➡️ For more examples, see [GoDoc](https://pkg.go.dev/github.com/shinagawa-web/strlistutils) or [examples](https://github.com/shinagawa-web/strlistutils/blob/main/strlistutils_example_test.go)

## API

The package intentionally keeps a small surface area:

| Function                                       | Behavior                                             |               |
| ---------------------------------------------- | ---------------------------------------------------- | ------------- |
| `RemoveDuplicates([]string) []string`          | Deduplicate while preserving first appearance order. |               |
| `TrimEach([]string) []string`                  | Trim leading/trailing spaces from each element.      |               |
| `FilterEmpty([]string) []string`               | Remove empty strings.                                |               |
| `Map([]string, func(string) string) []string`  | Apply a mapper to each element.                      |               |
| `Filter([]string, func(string) bool) []string` | Keep elements that satisfy the predicate.            |               |
| `Join([]string, string) string`                | Join with a separator.                               |               |
| `Reverse([]string) []string`                   | Reverse the order.                                   | ([GitHub][1]) |

## Performance and Reliability

- No allocations beyond what the operation requires (e.g., dedup uses a set and preserves order).
- Benchmarks and fuzz tests are included in the repository to keep regressions visible in CI. 

### Versioning Policy

This library aims to be stable. Minor releases may add functions; breaking changes will bump the major version. Pin a tag in go.mod when adopting.

### Real-world use

Typical use cases include CSV cleaning, environment variable parsing, and HTTP input normalization. A short write-up on the motivation and design is available on the project site/blog. 

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
