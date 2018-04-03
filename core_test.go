package strset

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type CoreSuite struct{}

var _ = Suite(&CoreSuite{})

func (s *CoreSuite) TestMake_empty(c *C) {
	set := Make()
	c.Assert(set.Len(), Equals, 0)
}