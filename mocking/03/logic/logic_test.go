package logic

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mattyhall/cb-advance-go-materials/mocking/03/client"
	mock_client "github.com/mattyhall/cb-advance-go-materials/mocking/03/client/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCountUoLUniversities(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
		cli = mock_client.NewMockClient(ctrl)
	)

	cli.EXPECT().GetUniversities().Return([]client.University{
		{Name: "Durham University"},
		{Name: "UCL, University of London"},
		{Name: "University of London"},
		{Name: "A University not of London"},
	}, nil)

	count, err := CountUoLUnis(cli)
	require.NoError(t, err)
	require.Equal(t, 2, count)
}

func TestCountUoLUniversitiesError(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
		cli = mock_client.NewMockClient(ctrl)
	)

	cli.EXPECT().GetUniversities().Return(nil, assert.AnError).Times(1)

	_, err := CountUoLUnis(cli)
	require.ErrorIs(t, err, assert.AnError)
}
