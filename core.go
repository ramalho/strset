package strset

// Set of strings.
type Set struct {
	store map[string]struct{}
}

// Make creates and returns a new Set.
func Make() Set {
	return Set{}
}

func (s Set) Len() int {
	return len(s.store)
}


