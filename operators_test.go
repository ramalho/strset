package strset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	set1 Set
	set2 Set
	want Set
}

var intersectionTestCases = []TestCase{
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

func TestIntersection(t *testing.T) {
	for _, tc := range intersectionTestCases {
		t.Run(fmt.Sprintf("%v.Intersection(%v) is %v", tc.set1, tc.set2, tc.want), func(t *testing.T) {
			got := tc.set1.Intersection(tc.set2)
			assert.Equal(t, tc.want, got)
		})
	}
}

var unionTestCases = []TestCase{
	{empty, empty, empty},
	{singleton, odd, odd},
	{singleton, even, Make("0", "1", "2", "4", "6", "8")},
	{even, odd, universe},
	{even, even, even},
	{universe, even, universe},
	{universe, empty, universe},
	{prime, fibonacci, Make("0", "1", "2", "3", "5", "7", "8")},
}

func TestUnion(t *testing.T) {
	for _, tc := range unionTestCases {
		t.Run(fmt.Sprintf("%v.Union(%v) is %v", tc.set1, tc.set2, tc.want), func(t *testing.T) {
			got := tc.set1.Union(tc.set2)
			assert.Equal(t, tc.want, got)
		})
	}
}

var differenceTestCases = []TestCase{
	{empty, empty, empty},
	{singleton, empty, singleton},
	{singleton, singleton, empty},
	{universe, even, odd},
	{prime, fibonacci, Make("7")},
	{fibonacci, prime, Make("0", "1", "8")},
}

func TestDifference(t *testing.T) {
	for _, tc := range differenceTestCases {
		t.Run(fmt.Sprintf("%v.Difference(%v) is %v", tc.set1, tc.set2, tc.want), func(t *testing.T) {
			got := tc.set1.Difference(tc.set2)
			assert.Equal(t, tc.want, got)
		})
	}
}

var symmetricDifferenceTestCases = []TestCase{
	{empty, empty, empty},
	{singleton, empty, singleton},
	{empty, singleton, singleton},
	{universe, even, odd},
	{odd, even, universe},
	{fibonacci, prime, Make("0", "1", "7", "8")},
}

func TestSymmetricDifference(t *testing.T) {
	for _, tc := range symmetricDifferenceTestCases {
		t.Run(fmt.Sprintf("%v.SymmetricDifference(%v) is %v", tc.set1, tc.set2, tc.want), func(t *testing.T) {
			got := tc.set1.SymmetricDifference(tc.set2)
			assert.Equal(t, tc.want, got)
		})
	}
}

func ExampleSet_SymmetricDifference() {
	s1 := MakeFromText("beta alpha delta gamma")
	s2 := MakeFromText("beta delta pi")
	fmt.Println(s1.SymmetricDifference(s2))
	// Output:
	// Set{alpha gamma pi}
}

func Example() {
	s1 := Make("red", "green", "blue", "yellow")
	s2 := MakeFromText("yellow green white")
	fmt.Println(s1.Intersection(s2))
	// Output: Set{green yellow}
}
