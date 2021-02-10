package hashset_test

import (
	"testing"

	hashset "github.com/kgoins/hashset/pkg"
	"github.com/stretchr/testify/require"
)

func TestContains(t *testing.T) {
	rq := require.New(t)
	set := hashset.NewStrHashset("val1", "VaR2")

	rq.True(set.Contains("val1"))
	rq.False(set.Contains("1"))
}

func TestContainsSubstr(t *testing.T) {
	rq := require.New(t)
	set := hashset.NewStrHashset("val1", "VaR2")

	rq.True(set.ContainsSubstr("1"))
	rq.False(set.ContainsSubstr("nothere"))

	rq.False(set.ContainsSubstr("var"))
	rq.True(set.ContainsSubstr("var", true))
}
