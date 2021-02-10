package hashset_test

import (
	"testing"

	hashset "github.com/kgoins/hashset/pkg"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	rq := require.New(t)

	emptySet := hashset.NewStrHashset()
	rq.NotNil(emptySet)
	rq.Equal(emptySet.Size(), 0)

	setWithVals := hashset.NewStrHashset("val1", "val2")
	rq.NotNil(setWithVals)
	rq.Equal(setWithVals.Size(), 2)

	var nilVals []string
	nilSet := hashset.NewStrHashset(nilVals...)
	rq.NotNil(nilSet)
	rq.Equal(nilSet.Size(), 0)
}

func TestClear(t *testing.T) {
	rq := require.New(t)

	set := hashset.NewStrHashset("val1", "val2")
	set.Clear()

	rq.Equal(set.Size(), 0)
}

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
