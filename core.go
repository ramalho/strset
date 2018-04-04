// Package strset provides a Set type for string elements.
package strset

/* Note: The only methods that change a Set after it is made
   are in updaters.go
*/

import (
	"bytes"
	"sort"
	"strings"
)

// Set of strings.
type Set struct {
	store map[string]struct{}
}

// Make creates and returns a new Set.
func Make(elems ...string) Set {
	s := Set{}
	s.store = make(map[string]struct{})
	for _, elem := range elems {
		s.store[elem] = struct{}{}
	}
	return s
}

// MakeFromText creates and returns a new Set from
// a string of elements separated by whitespace.
func MakeFromText(text string) Set {
	return Make(strings.Fields(text)...)
}

// Len reports the number of elements in the set.
func (s Set) Len() int {
	return len(s.store)
}

// Has reports whether set contains the element.
// Math: S ∋ e.
func (s Set) Has(elem string) bool {
	_, found := s.store[elem]
	return found
}

// HasAll reports whether s contains all the given elements
func (s Set) HasAll(elems ...string) bool {
	for _, elem := range elems {
		if !s.Has(elem) {
			return false
		}
	}
	return true
}

// Elems returns a new slice with the elements of s.
// The order of the elements is undefined.
func (s Set) Elems() []string {
	var elems []string
	for elem := range s.store {
		elems = append(elems, elem)
	}
	return elems
}

// String returns a string representation of s with
// elements in lexicographic order.
func (s Set) String() string {
	elems := s.Elems()
	sort.Strings(elems)
	var buf bytes.Buffer
	buf.WriteString("Set{")
	buf.WriteString(strings.Join(elems, " "))
	buf.WriteByte('}')
	return buf.String()
}

// allIn reports whether all elements of s exist in other.
func (s Set) allIn(other Set) bool {
	for elem := range s.store {
		if _, found := other.store[elem]; !found {
			return false
		}
	}
	return true
}

// Equal reports whether set is equal to other
func (s Set) Equal(other Set) bool {
	return len(s.store) == len(other.store) && s.allIn(other)
}

// Copy returns a new Set: a copy of s.
func (s Set) Copy() Set {
	res := Make()
	for elem := range s.store {
		res.Add(elem)
	}
	return res
}
