package strset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsetOf(t *testing.T) {
	testCases := []struct {
		set1 Set
		set2 Set
		want bool
	}{
		{empty, empty, true},
		{singleton, empty, false},
		{empty, singleton, true},
		{universe, even, false},
		{even, universe, true},
		{fibonacci, prime, false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v.SubsetOf(%v) is %v", tc.set1, tc.set2, tc.want), func(t *testing.T) {
			got := tc.set1.SubsetOf(tc.set2)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSupersetOf(t *testing.T) {
	testCases := []struct {
		set1 Set
		set2 Set
		want bool
	}{
		{empty, empty, true},
		{singleton, empty, true},
		{empty, singleton, false},
		{universe, even, true},
		{even, universe, false},
		{fibonacci, prime, false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v.SupersetOf(%v) is %v", tc.set1, tc.set2, tc.want), func(t *testing.T) {
			got := tc.set1.SupersetOf(tc.set2)
			assert.Equal(t, tc.want, got)
		})
	}
}
