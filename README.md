# utils

utils based on generics

- Contains returns true if element is in array
- Keys returns keys of map
- Filter returns elements of array that match predicate
- MapSlice returns map of elements
- Reduce returns reduced value of elements
- Uniq removes duplications in slice


```go
import "github.com/jaigouk/utils"

elems := []string{"a", "b", "c"}
utils.Contains(elems, "b") // bool

m := map[string]string{"a": "b", "c": "d"}
utils.Keys(m) // []string

slice := []string{"a", "b", "c"}
f := func(s string) bool {
  return s == "b"
}
n := utils.Filter(slice, f) // []string{"b"}

slice := []int{3, 2, 1}
utils.SortSlice(slice) // []int{1, 2, 3}

mapped := utils.MapSlice(slice, func(i int) int {
  return i * 2
}) // []int{6, 4, 2}

slice := []int{1, 2, 3}
sum := mapped := utils.MapSlice(slice, func(i int) int {
.Reduce(slice, func(i int, j int) int {
  return i + j
}, 0) // returns 6
```
