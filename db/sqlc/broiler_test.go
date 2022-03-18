package db

import (
	"context"
	"github.com/ngenohkevin/flockmanager_backend/db/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomBroiler(t *testing.T) Broiler {
	arg := util.RandomTitle()

	broiler, err := testQueries.CreateBroiler(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, broiler)

	require.Equal(t, arg, broiler.Title)
	require.NotZero(t, broiler.ID)
	require.NotZero(t, broiler.CreatedAt)

	return broiler
}

func TestCreateBroiler(t *testing.T) {
	CreateRandomBroiler(t)
}

func TestGetBroiler(t *testing.T) {
	broiler1 := CreateRandomBroiler(t)
	broiler2, err := testQueries.GetBroiler(context.Background(), broiler1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, broiler2)

	require.Equal(t, broiler1.ID, broiler2.ID)
	require.Equal(t, broiler1.Title, broiler2.Title)
	require.WithinDuration(t, broiler1.CreatedAt, broiler2.CreatedAt, time.Second)
}

func TestDeleteBroiler(t *testing.T) {
	broiler1 := CreateRandomBroiler(t)
	err := testQueries.DeleteBroiler(context.Background(), broiler1.ID)
	require.NoError(t, err)

	broiler2, err := testQueries.GetBroiler(context.Background(), broiler1.ID)
	require.Error(t, err)
	require.Empty(t, broiler2)
}

func TestListBroiler(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomBroiler(t)
	}
	arg := ListBroilerParams{
		Limit:  5,
		Offset: 5,
	}
	broiler, err := testQueries.ListBroiler(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, broiler, 5)

	for _, broiler := range broiler {
		require.NotEmpty(t, broiler)
	}
}
