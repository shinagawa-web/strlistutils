# strlistutils

![Test](https://github.com/shinagawa-web/strlistutils/actions/workflows/test.yml/badge.svg)


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
go get github.com/shinagawa-web/strlistutils@v0.1.0
```

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

## Roadmap

- [x] RemoveDuplicates
- [x] TrimEach
- [x] FilterEmpty
- [ ] Contains
- [ ] IndexOf

## License

MIT
