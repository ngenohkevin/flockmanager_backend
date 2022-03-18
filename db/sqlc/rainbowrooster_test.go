package db

import (
	"context"
	"github.com/ngenohkevin/flockmanager_backend/db/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomRainbowRooster(t *testing.T) Rainbowrooster {
	arg := util.RandomTitle()

	rainbow, err := testQueries.CreateRainbowRooster(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, rainbow)

	require.Equal(t, arg, rainbow.Title)
	require.NotZero(t, rainbow.ID)
	require.NotZero(t, rainbow.CreatedAt)

	return rainbow
}

func TestCreateRainbowRooster(t *testing.T) {
	CreateRandomRainbowRooster(t)
}

func TestGetRainbowRooster(t *testing.T) {
	rainbow1 := CreateRandomRainbowRooster(t)
	rainbow2, err := testQueries.GetRainbowRooster(context.Background(), rainbow1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, rainbow2)

	require.Equal(t, rainbow1.ID, rainbow2.ID)
	require.Equal(t, rainbow1.Title, rainbow2.Title)
	require.WithinDuration(t, rainbow1.CreatedAt, rainbow2.CreatedAt, time.Second)
}

func TestDeleteRainbowRooster(t *testing.T) {
	rainbow1 := CreateRandomRainbowRooster(t)
	err := testQueries.DeleteRainbowRooster(context.Background(), rainbow1.ID)
	require.NoError(t, err)

	rainbow2, err := testQueries.GetRainbowRooster(context.Background(), rainbow1.ID)
	require.Error(t, err)
	require.Empty(t, rainbow2)
}

func TestListRainbowRooster(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomRainbowRooster(t)
	}
	arg := ListRainbowRoosterParams{
		Limit:  5,
		Offset: 5,
	}
	rainbow, err := testQueries.ListRainbowRooster(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, rainbow, 5)

	for _, rainbow := range rainbow {
		require.NotEmpty(t, rainbow)
	}
}
