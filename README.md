# strset

[![GoDoc](https://godoc.org/github.com/standupdev/strset?status.svg)](https://godoc.org/github.com/standupdev/strset)

Full-featured Go `Set` type for string elements.

```golang
func Example() {
	s1 := Make("red", "green", "blue", "yellow")
	s2 := MakeFromText("yellow green white")
	fmt.Println(s1.Intersection(s2))
	// Output: Set{green yellow}
}
```

Main features:

* `Make` builds a set from zero or more strings (or `[]string...`).
* `MakeFromText` builds a set from a single string with elements separated by whitespace.
* `String` method returns elements in ascending order.

Happy hacking!
