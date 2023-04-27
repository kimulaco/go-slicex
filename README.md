# go-slicex

Utility functions module for handling slice by go.

## Usage

```go
// IndexOf
i = slicex.IndexOf([]string{"1", "2", "3"}, 2) // 1
i = slicex.IndexOf([]string{"1", "2", "3"}, 5) // -1

// Contains
b = slicex.Contains([]string{"1", "2", "3"}, 2) // true
b = slicex.Contains([]string{"1", "2", "3"}, 5) // false

// RemoveIndex
s = slicex.RemoveIndex([]string{"1", "2", "3"}, 2) // []string{"1", "2"}

// RemoveValue
s = slicex.RemoveValue([]string{"1", "2", "3"}, "2") // []string{"1", "3"}

// ForEach
slicex.ForEach([]string{"1", "2", "3"}, func (v string, i, int) {
    // v: "1", "2", "3"
    // i: 0, 1, 2
})

// Filter
s = slicex.Filter([]string{"1", "2", "3"}, func (v string, i, int) bool {
    return v != "2" || i > 0
})
// []string{"3"}

// Map
s = slicex.Map([]string{"1", "2", "3"}, func (v string, i, int) bool {
    return v + v
})
// []string{"11", "22", "33"}
```
