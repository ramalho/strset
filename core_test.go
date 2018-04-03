package strset

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
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
		{[]string{"a", "b"}, "c", false},
		{[]string{"a", "b"}, "b", true},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%q in %v is %v", tc.elems, tc.needle, tc.want), func(t *testing.T) {
			s := Make(tc.elems...)
			got := s.Has(tc.needle)
			assert.Equal(t, tc.want, got)
		})
	}
}
