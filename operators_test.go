package strset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	universe  = Make("0", "1", "2", "3", "4", "5", "6", "7", "8", "9")
	even      = Make("0",      "2",      "4",      "6",      "8")
	odd       = Make(     "1",      "3",      "5",      "7",      "9")
	prime     = Make(          "2", "3",      "5",      "7")
	fibonacci = Make("0", "1", "2", "3",      "5",           "8")
	singleton = Make(     "1")
	empty     = Make()
)

func TestIntersection(t *testing.T) {
	testCases := []struct {
		set1 Set
		set2 Set
		want Set
	}{
		{empty, empty, empty},
		{singleton, even, empty},
		{singleton, odd, singleton},
		{even, odd, empty},
		{even, even, even},
		{universe, even, even},
		{universe, empty, empty},
		{prime, fibonacci, Make("2", "3", "5")},
		{fibonacci, prime, Make("2", "3", "5")},

	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v.Intersection(%v) is %v", tc.set1, tc.set2, tc.want), func(t *testing.T) {
			got := tc.set1.Intersection(tc.set2)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestUnion(t *testing.T) {
	testCases := []struct {
		set1 Set
		set2 Set
		want Set
	}{
		{empty, empty, empty},
		{singleton, odd, odd},
		{singleton, even, Make("0","1", "2", "4", "6", "8")},
		{even, odd, universe},
		{even, even, even},
		{universe, even, universe},
		{universe, empty, universe},
		{prime, fibonacci, Make("0", "1", "2", "3", "5", "7", "8")},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v.Union(%v) is %v", tc.set1, tc.set2, tc.want), func(t *testing.T) {
			got := tc.set1.Union(tc.set2)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestDifference(t *testing.T) {
	testCases := []struct {
		set1 Set
		set2 Set
		want Set
	}{
		{empty, empty, empty},
		{singleton, empty, singleton},
		{singleton, singleton, empty},
		{universe, even, odd},
		{prime, fibonacci, Make("7")},
		{fibonacci, prime, Make("0", "1", "8")},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v.Difference(%v) is %v", tc.set1, tc.set2, tc.want), func(t *testing.T) {
			got := tc.set1.Difference(tc.set2)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSymmetricDifference(t *testing.T) {
	testCases := []struct {
		set1 Set
		set2 Set
		want Set
	}{
		{empty, empty, empty},
		{singleton, empty, singleton},
		{empty, singleton, singleton},
		{universe, even, odd},
		{odd, even, universe},
		{fibonacci, prime, Make("0", "1", "7", "8")},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v.SymmetricDifference(%v) is %v", tc.set1, tc.set2, tc.want), func(t *testing.T) {
			got := tc.set1.SymmetricDifference(tc.set2)
			assert.Equal(t, tc.want, got)
		})
	}
}
