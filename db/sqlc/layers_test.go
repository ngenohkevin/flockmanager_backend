package db

import (
	"context"
	"github.com/ngenohkevin/flockmanager_backend/db/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomLayers(t *testing.T) Layer {
	arg := util.RandomTitle()

	layers, err := testQueries.CreateLayers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, layers)

	require.Equal(t, arg, layers.Title)
	require.NotZero(t, layers.ID)
	require.NotZero(t, layers.CreatedAt)

	return layers
}

func TestCreateLayers(t *testing.T) {
	CreateRandomLayers(t)
}

func TestGetLayers(t *testing.T) {
	layers1 := CreateRandomLayers(t)
	layers2, err := testQueries.GetLayers(context.Background(), layers1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, layers2)

	require.Equal(t, layers1.ID, layers2.ID)
	require.Equal(t, layers1.Title, layers2.Title)
	require.WithinDuration(t, layers1.CreatedAt, layers2.CreatedAt, time.Second)
}

func TestDeleteLayers(t *testing.T) {
	layers1 := CreateRandomLayers(t)
	err := testQueries.DeleteLayers(context.Background(), layers1.ID)
	require.NoError(t, err)

	layers2, err := testQueries.GetLayers(context.Background(), layers1.ID)
	require.Error(t, err)
	require.Empty(t, layers2)
}

func TestListLayers(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomLayers(t)
	}
	arg := ListLayersParams{
		Limit:  5,
		Offset: 5,
	}
	layers, err := testQueries.ListLayers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, layers, 5)

	for _, layers := range layers {
		require.NotEmpty(t, layers)
	}
}
