package strset

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
	s.AddAll(elems...)
	return s
}

// Add adds element to set.
func (s Set) Add(elem string) {
	s.store[elem] = struct{}{}
}

// AddAll adds elements to set.
func (s Set) AddAll(elems ...string) {
	for _, elem := range elems {
		s.Add(elem)
	}
}

// Len reports the number of elements in the set.
func (s Set) Len() int {
	return len(s.store)
}

// Has reports whether set contains the element.
func (s Set) Has(elem string) bool {
	_, found := s.store[elem]
	return found
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




