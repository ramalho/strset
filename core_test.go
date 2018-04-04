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
		set    Set
		needle string
		want   bool
	}{
		{Make(), "a", false},
		{Make("a"), "a", true},
		{Make("a", "b"), "b", true},
		{Make("a", "b"), "c", false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%q in %v is %v", tc.set, tc.needle, tc.want), func(t *testing.T) {
			got := tc.set.Has(tc.needle)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestString(t *testing.T) {
	testCases := []struct {
		set  Set
		want string
	}{
		{Make(), `Set{}`},
		{Make("a"), `Set{a}`},
		{Make("b", "a"), `Set{a b}`},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v gets %q", tc.set, tc.want), func(t *testing.T) {
			got := fmt.Sprint(tc.set)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestEqual(t *testing.T) {
	testCases := []struct {
		set1 Set
		set2 Set
		want bool
	}{
		{Make(), Make(), true},
		{Make("a"), Make(), false},
		{Make(), Make("b"), false},
		{Make("a"), Make("a"), true},
		{Make("a"), Make("b"), false},
		{Make("a", "b"), Make("a", "b"), true},
		{Make("a", "b"), Make("b", "a"), true},
		{Make("a", "b"), Make("a", "b", "c"), false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v eq %v is %v", tc.set1, tc.set2, tc.want), func(t *testing.T) {
			got := tc.set1.Equal(tc.set2)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		s1 Set
		out string
		s2 Set
	}{
		{Make(), "a", Make()},
		{Make("a"), "a",  Make()},
		{Make("a"), "b", Make("a")},
		{Make("a", "b"), "a", Make("b")},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v remove %v remain %v", tc.s1, tc.out, tc.s2),
			func(t *testing.T) {
				tc.s1.Remove(tc.out)
				assert.True(t, tc.s1.Equal(tc.s2))
			})
	}
}

