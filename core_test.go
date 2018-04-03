package strset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	s := Make()
	assert.Equal(t, 0, s.Len())
}
