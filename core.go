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

func (s Set) Len() int {
	return len(s.store)
}


