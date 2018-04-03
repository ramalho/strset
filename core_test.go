package strset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMake_empty(t *testing.T) {
	s := Make()
	assert.Equal(t, 0, s.Len())
}

func TestMake(t *testing.T) {
	testCases := []struct {
		elems []string
		want int
	}{
		{[]string{}, 0},
		{[]string{"a"}, 1},
		{[]string{"a", "b"}, 2},
		{[]string{"a", "b", "a"}, 2},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v gets %d", tc.elems, tc.want), func(t *testing.T) {
			s := Make(tc.elems...)
			assert.Equal(t, tc.want, s.Len())
		})
	}
}

func TestHas(t *testing.T) {
	testCases := []struct {
		elems []string
		needle string
		want bool
	}{
		{[]string{}, "a", false},
		{[]string{"a"}, "a", true},
		{[]string{"a", "b"}, "b", true},
		{[]string{"a", "b"}, "c", false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%q in %v is %v", tc.elems, tc.needle, tc.want), func(t *testing.T) {
			s := Make(tc.elems...)
			got := s.Has(tc.needle)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestString(t *testing.T) {
	testCases := []struct {
		elems []string
		want string
	}{
		{[]string{}, `Set{}`},
		{[]string{"a"}, `Set{a}`},
		{[]string{"b", "a"}, `Set{a b}`},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v gets %q", tc.elems, tc.want), func(t *testing.T) {
			s := Make(tc.elems...)
			got := fmt.Sprint(s)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestEqual(t *testing.T) {
	testCases := []struct {
		elems1 []string
		elems2 []string
		want bool
	}{
		{[]string{}, []string{}, true},
		{[]string{"a"}, []string{}, false},
		{[]string{}, []string{"b"}, false},
		{[]string{"a"}, []string{"a"}, true},
		{[]string{"a"}, []string{"b"}, false},
		{[]string{"a", "b"}, []string{"a", "b"}, true},
		{[]string{"a", "b"}, []string{"b", "a"}, true},
		{[]string{"a", "b"}, []string{"a", "b", "c"}, false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v eq %v is %v", tc.elems1, tc.elems2, tc.want), func(t *testing.T) {
			s1 := Make(tc.elems1...)
			s2 := Make(tc.elems2...)
			got := s1.Equal(s2)
			assert.Equal(t, tc.want, got)
		})
	}
}


