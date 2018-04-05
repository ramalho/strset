# strset

[![GoDoc](https://godoc.org/github.com/standupdev/strset?status.svg)](https://godoc.org/github.com/standupdev/strset)

Full-featured Go `Set` type for `string` elements.

```golang
func Example() {
	s1 := Make("red", "green", "blue", "yellow")
	s2 := MakeFromText("yellow green white")
	fmt.Println(s1.Intersection(s2))
	// Output: Set{green yellow}
}
```

Some features of the `strset.Set` type:

* `Make` builds a set from zero or more strings (or `[]string...`).
* `MakeFromText` builds a set from a single string with elements separated by whitespace.
* `String` method returns elements in ascending order.
* Methods returning new sets: intersection, union, difference, symmetric difference.
* Methods updating receiver in-place for each operation above.
* `Pop` method to retrieve and delete one unspecified element.
* Need an immutable set? Just remove `updaters.go` from the build.
* 100% test coverage.
* Not thread-safe.

Happy hacking!
