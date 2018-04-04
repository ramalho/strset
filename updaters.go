package strset

// IntersectionUpdate keeps in s only elements that are in s AND in other.
// Math: S ∩ Z.
func (s Set) IntersectionUpdate(other Set) {
	for elem := range s.store {
		if !other.Has(elem) {
			s.Remove(elem)
		}
	}
}

// UnionUpdate keeps/puts in s all elements that are in s OR in other.
// Math: S ∪ Z.
func (s Set) UnionUpdate(other Set) {
	for elem := range other.store {
		s.Add(elem)
	}
}


// DifferenceUpdate removes from s the elements of other.
// Math: S \ Z.
func (s Set) DifferenceUpdate(other Set) {
	for elem := range other.store {
		s.Remove(elem)
	}
}

// SymmetricDifferenceUpdate keeps in s only elements that
// are in either set but not on both. Think boolean XOR.
// Math: S ∆ Z.
func (s Set) SymmetricDifferenceUpdate(other Set) {
	common := s.Intersection(other)
	s.UnionUpdate(other)
	s.DifferenceUpdate(common)
}

