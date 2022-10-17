package unis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountUoLUniversities(t *testing.T) {
	cli := MockClient{results: []University{
		{Name: "Durham University"},
		{Name: "UCL, University of London"},
		{Name: "University of London"},
		{Name: "A University not of London"},
	}}

	count, err := CountUoLUnis(&cli)
	require.NoError(t, err)
	require.Equal(t, 2, count)
}

type Multiple interface {
	FindMultiple(a, b int) int
}
