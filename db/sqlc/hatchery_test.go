package db

import (
	"context"
	"github.com/ngenohkevin/flockmanager_backend/db/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomHatchery(t *testing.T, kuroiler Kuroiler) Hatchery {
	args := CreateHatcheryParams{
		HatcheryID:  kuroiler.ID,
		Infertile:   util.RandomNums(),
		Early:       util.RandomNums(),
		Middle:      util.RandomNums(),
		Late:        util.RandomNums(),
		DeadChicks:  util.RandomNums(),
		AliveChicks: util.RandomNums(),
	}
	hatchery, err := testQueries.CreateHatchery(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, hatchery)

	require.Equal(t, args.HatcheryID, hatchery.HatcheryID)
	require.Equal(t, args.Infertile, hatchery.Infertile)
	require.Equal(t, args.Early, hatchery.Early)
	require.Equal(t, args.Middle, hatchery.Middle)
	require.Equal(t, args.Late, hatchery.Late)
	require.Equal(t, args.DeadChicks, hatchery.DeadChicks)
	require.Equal(t, args.AliveChicks, hatchery.AliveChicks)

	require.NotZero(t, hatchery.ID)
	require.NotZero(t, hatchery.CreatedAt)

	return hatchery
}

func TestGetHatchery(t *testing.T) {
	kuroiler := CreateRandomKuroiler(t)
	hatchery1 := CreateRandomHatchery(t, kuroiler)
	hatchery2, err := testQueries.GetHatchery(context.Background(), hatchery1.ID)
	require.NoError(t, err)

	require.Equal(t, hatchery1.ID, hatchery2.ID)
	require.Equal(t, hatchery1.HatcheryID, hatchery2.HatcheryID)
	require.Equal(t, hatchery1.Infertile, hatchery2.Infertile)
	require.Equal(t, hatchery1.Early, hatchery2.Early)
	require.Equal(t, hatchery1.Middle, hatchery2.Middle)
	require.Equal(t, hatchery1.Late, hatchery2.Late)
	require.Equal(t, hatchery1.DeadChicks, hatchery2.DeadChicks)
	require.Equal(t, hatchery1.AliveChicks, hatchery2.AliveChicks)

	require.WithinDuration(t, hatchery1.CreatedAt, hatchery2.CreatedAt, time.Second)
}

func TestListHatchery(t *testing.T) {
	hat := CreateRandomKuroiler(t)
	for i := 0; i < 10; i++ {
		CreateRandomHatchery(t, hat)
	}
	arg := ListHatcheryParams{
		HatcheryID: hat.ID,
		Limit:      5,
		Offset:     5,
	}
	hatchery, err := testQueries.ListHatchery(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, hatchery, 5)

	for _, hatchery := range hatchery {
		require.NotEmpty(t, hatchery)
		require.Equal(t, arg.HatcheryID, hatchery.HatcheryID)
	}
}
