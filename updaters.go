package strset

/* Note: The only methods that change a Set after it is made
   are in this source file.
*/


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

// Remove removes element from the set, if it is present.
func (s Set) Remove(elem string) {
	delete(s.store, elem)
}

// RemoveAll removes elements from the set, if they are present.
func (s Set) RemoveAll(elems ...string) {
	for _, elem := range elems {
		s.Remove(elem)
	}
}

// Pop tries to return some element of s, deleting it. If there was an element,
// the pair (element, true) is returned. Otherwise, the result is ("", false).
func (s Set) Pop() (elem string, found bool) {
	for elem = range s.store {
		delete(s.store, elem)
		return elem, true
	}
	return "", false
}

// Clear removes all elements from s.
func (s *Set) Clear() {
	s.store = make(map[string]struct{})
}

/* Set operations that change the receiver */

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

