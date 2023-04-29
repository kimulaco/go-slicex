# go-slicex

[![Test](https://github.com/kimulaco/go-slicex/actions/workflows/test.yml/badge.svg)](https://github.com/kimulaco/go-slicex/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/kimulaco/go-slicex/branch/main/graph/badge.svg?token=33P1XO6YK3)](https://codecov.io/gh/kimulaco/go-slicex)
[![License MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

Utility functions module for handling slice by go.

## Install

```bash
go get github.com/kimulaco/go-slicex
```

## Usage

```go
// IndexOf
i = slicex.IndexOf([]string{"1", "2", "3"}, "2") // 1
i = slicex.IndexOf([]string{"1", "2", "3"}, "5") // -1

// Contains
b = slicex.Contains([]string{"1", "2", "3"}, "2") // true
b = slicex.Contains([]string{"1", "2", "3"}, "5") // false

// RemoveIndex
s = slicex.RemoveIndex([]string{"1", "2", "3"}, 2) // []string{"1", "2"}

// RemoveValue
s = slicex.RemoveValue([]string{"1", "2", "3"}, "2") // []string{"1", "3"}

// RemoveValueAll
s = slicex.RemoveValueAll([]string{"1", "2", "2", "3", "2"}, "2") // []string{"1", "3"}

// ForEach
slicex.ForEach([]string{"1", "2", "3"}, func(v string, i int) {
    // v: "1", "2", "3"
    // i: 0, 1, 2
})

// Filter
s = slicex.Filter([]string{"1", "2", "3"}, func(v string, i int) bool {
    return v != "2" || i > 0
})
// []string{"3"}

// Map
s = slicex.Map([]string{"1", "2", "3"}, func(v string, i int) string {
    return v + v
})
// []string{"11", "22", "33"}
```
