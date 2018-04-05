package strset

// SubsetOf reports whether s is a subset of other
// Math: S ⊆ Z.
func (s Set) SubsetOf(other Set) bool {
	if s.Len() > other.Len() {
		return false
	}
	return s.allIn(other)
}

// SupersetOf reports whether s is a superset of other
// Math: S ⊇ Z.
func (s Set) SupersetOf(other Set) bool {
	return other.SubsetOf(s)
}
