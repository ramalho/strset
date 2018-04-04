package strset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersectionUpdate(t *testing.T) {
	for _, tc := range intersectionTestCases {
		t.Run(fmt.Sprintf("%v.IntersectionUpdate(%v) is %v", tc.set1, tc.set2, tc.want), func(t *testing.T) {
			got := tc.set1.Copy()
			got.IntersectionUpdate(tc.set2)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestUnionUpdate(t *testing.T) {
	for _, tc := range unionTestCases {
		t.Run(fmt.Sprintf("%v.UnionUpdate(%v) is %v", tc.set1, tc.set2, tc.want), func(t *testing.T) {
			got := tc.set1.Copy()
			got.UnionUpdate(tc.set2)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestDifferenceUpdate(t *testing.T) {
	for _, tc := range differenceTestCases {
		t.Run(fmt.Sprintf("%v.DifferenceUpdate(%v) is %v", tc.set1, tc.set2, tc.want), func(t *testing.T) {
			got := tc.set1.Copy()
			got.DifferenceUpdate(tc.set2)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSymmetricDifferenceUpdate(t *testing.T) {
	for _, tc := range symmetricDifferenceTestCases {
		t.Run(fmt.Sprintf("%v.SymmetricDifferenceUpdate(%v) is %v", tc.set1, tc.set2, tc.want), func(t *testing.T) {
			got := tc.set1.Copy()
			got.SymmetricDifferenceUpdate(tc.set2)
			assert.Equal(t, tc.want, got)
		})
	}
}
