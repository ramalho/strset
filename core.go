package strset

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

// Len reports the number of elements in the set.
func (s Set) Len() int {
	return len(s.store)
}

// Has reports whether set contains the element.
func (s Set) Has(elem string) bool {
	_, found := s.store[elem]
	return found
}



