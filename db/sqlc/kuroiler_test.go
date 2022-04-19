package db

import (
	"context"
	"github.com/ngenohkevin/flockmanager_backend/db/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomKuroiler(t *testing.T) Kuroiler {
	arg := CreateKuroilerParams{
		Title: util.RandomTitle(),
		House: util.RandomHouse(),
	}

	kuroiler, err := testQueries.CreateKuroiler(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, kuroiler)

	require.Equal(t, arg.Title, kuroiler.Title)
	require.Equal(t, arg.House, kuroiler.House)

	require.NotZero(t, kuroiler.ID)
	require.NotZero(t, kuroiler.CreatedAt)

	return kuroiler
}

func TestCreateKuroiler(t *testing.T) {
	CreateRandomKuroiler(t)
}

func TestGetKuroiler(t *testing.T) {
	kuroiler1 := CreateRandomKuroiler(t)
	kuroiler2, err := testQueries.GetKuroiler(context.Background(), kuroiler1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, kuroiler2)

	require.Equal(t, kuroiler1.ID, kuroiler2.ID)
	require.Equal(t, kuroiler1.Title, kuroiler2.Title)
	require.Equal(t, kuroiler1.House, kuroiler2.House)
	require.WithinDuration(t, kuroiler1.CreatedAt, kuroiler2.CreatedAt, time.Second)
}

func TestDeleteKuroiler(t *testing.T) {
	kuroiler1 := CreateRandomKuroiler(t)
	err := testQueries.DeleteKuroiler(context.Background(), kuroiler1.ID)
	require.NoError(t, err)

	kuroiler2, err := testQueries.GetKuroiler(context.Background(), kuroiler1.ID)
	require.Error(t, err)
	require.Empty(t, kuroiler2)
}

func TestListKuroiler(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomKuroiler(t)
	}
	arg := ListKuroilerParams{
		Limit:  5,
		Offset: 5,
	}
	kuroiler, err := testQueries.ListKuroiler(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, kuroiler, 5)

	for _, kuroiler := range kuroiler {
		require.NotEmpty(t, kuroiler)
	}
}
