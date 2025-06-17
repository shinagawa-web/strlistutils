# strlistutils

**strlistutils** is a small utility library for working with `[]string` in Go.  
It provides common operations like removing duplicates, trimming, filtering, and more.

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

## Roadmap

- [x] RemoveDuplicates
- [ ] TrimEach
- [ ] FilterEmpty
- [ ] Contains
- [ ] IndexOf

## License

MIT
