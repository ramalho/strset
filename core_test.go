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
		elems   []string
		wantLen int
	}{
		{[]string{}, 0},
		{[]string{"a"}, 1},
		{[]string{"a", "b"}, 2},
		{[]string{"a", "b", "a"}, 2},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v gets %d", tc.elems, tc.wantLen), func(t *testing.T) {
			s := Make(tc.elems...)
			assert.Equal(t, tc.wantLen, s.Len())
		})
	}
}

func ExampleMake() {
	w := []string{"beta", "alpha", "gamma", "beta"}
	s := Make(w...)
	fmt.Println(s)
	// Output: Set{alpha beta gamma}
}

func TestContains(t *testing.T) {
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
			got := tc.set.Contains(tc.needle)
			assert.Equal(t, tc.want, got)
		})
	}
}

var (
	universe  = MakeFromText("0 1 2 3 4 5 6 7 8 9")
	even      = MakeFromText("0   2   4   6   8  ")
	odd       = MakeFromText("  1   3   5   7   9")
	prime     = MakeFromText("    2 3   5   7    ")
	fibonacci = MakeFromText("0 1 2 3   5     8  ")
	singleton = MakeFromText("  1                ")
	empty     = Make()
)

func TestContainsAll(t *testing.T) {
	testCases := []struct {
		set   Set
		slice []string
		want  bool
	}{
		{empty, empty.ToSlice(), true},
		{singleton, empty.ToSlice(), true},
		{empty, singleton.ToSlice(), false},
		{universe, even.ToSlice(), true},
		{even, universe.ToSlice(), false},
		{fibonacci, prime.ToSlice(), false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v.ContainsAll(%v...) is %v", tc.set, tc.slice, tc.want), func(t *testing.T) {
			got := tc.set.ContainsAll(tc.slice...)
			assert.Equal(t, tc.want, got)
		})
	}
}

func ExampleSet_ContainsAll() {
	s := MakeFromText("alpha beta gamma")
	query := []string{"gamma", "beta"}
	fmt.Println(s.ContainsAll(query...))
	// Output: true
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

func TestCopy(t *testing.T) {
	testCases := []Set{
		Make(),
		Make("a"),
		Make("a", "b"),
	}
	for _, set := range testCases {
		t.Run(fmt.Sprintf("%v.Copy()", set), func(t *testing.T) {
			clone := set.Copy()
			assert.True(t, set.Equal(clone))
			set.store["zzz"] = struct{}{}
			assert.False(t, set.Equal(clone))
		})
	}
}

func TestMakeFromText(t *testing.T) {
	testCases := []struct {
		text string
		want Set
	}{
		{"", Make()},
		{"  ", Make()},
		{" a ", Make("a")},
		{"  b a ", Make("a", "b")},
		{"a b a", Make("a", "b")},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v makes %v", tc.text, tc.want), func(t *testing.T) {
			got := MakeFromText(tc.text)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestChannel(t *testing.T) {
	testCases := []struct {
		set  Set
		want []string
	}{
		{Make(), []string{}},
		{Make("a"), []string{"a"}},
		{Make("b", "a"), []string{"a", "b"}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v yields %v", tc.set, tc.want), func(t *testing.T) {
			got := []string{}
			for elem := range tc.set.Channel() {
				got = append(got, elem)
			}
			assert.ElementsMatch(t, tc.want, got)
		})
	}
}

func ExampleSet_Channel() {
	set := MakeFromText("beta alpha delta gamma")
	// iteration order over underlying map is undefined
	for elem := range set.Channel() {
		fmt.Println(elem)
	}
	// Unordered output:
	// alpha
	// beta
	// delta
	// gamma
}
