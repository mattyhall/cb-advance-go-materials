package scores

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxBreak(t *testing.T) {
	require.Equal(t, 147, MaxBreak())
}
