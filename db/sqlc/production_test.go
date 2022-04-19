package db

import (
	"context"
	"github.com/ngenohkevin/flockmanager_backend/db/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomProduction(t *testing.T, kuroiler Kuroiler) Production {
	args := CreateProductionParams{
		ProductionID: kuroiler.ID,
		Eggs:         util.RandomNums(),
		Dirty:        util.RandomNums(),
		WrongShape:   util.RandomNums(),
		WeakShell:    util.RandomNums(),
		Damaged:      util.RandomNums(),
		HatchingEggs: util.RandomNums(),
	}
	production, err := testQueries.CreateProduction(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, production)

	require.Equal(t, args.ProductionID, production.ProductionID)
	require.Equal(t, args.Eggs, production.Eggs)
	require.Equal(t, args.Dirty, production.Dirty)
	require.Equal(t, args.WrongShape, production.WrongShape)
	require.Equal(t, args.WeakShell, production.WeakShell)
	require.Equal(t, args.Damaged, production.Damaged)
	require.Equal(t, args.HatchingEggs, production.HatchingEggs)

	require.NotZero(t, production.ID)
	require.NotZero(t, production.CreatedAt)

	return production
}

//func TestCreateProduction(t *testing.T) {
//	CreateRandomProduction(t)
//}

func TestGetProduction(t *testing.T) {
	kuroiler := CreateRandomKuroiler(t)
	production1 := CreateRandomProduction(t, kuroiler)
	production2, err := testQueries.GetProduction(context.Background(), production1.ID)
	require.NoError(t, err)

	require.Equal(t, production1.ID, production2.ID)
	require.Equal(t, production1.ProductionID, production2.ProductionID)
	require.Equal(t, production1.Eggs, production2.Eggs)
	require.Equal(t, production1.Dirty, production2.Dirty)
	require.Equal(t, production1.WrongShape, production2.WrongShape)
	require.Equal(t, production1.WeakShell, production2.WeakShell)
	require.Equal(t, production1.Damaged, production2.Damaged)
	require.Equal(t, production1.HatchingEggs, production2.HatchingEggs)

	require.WithinDuration(t, production1.CreatedAt, production2.CreatedAt, time.Second)
}

func TestListProduction(t *testing.T) {
	prod := CreateRandomKuroiler(t)
	for i := 0; i < 10; i++ {
		CreateRandomProduction(t, prod)
	}
	arg := ListProductionParams{
		ProductionID: prod.ID,
		Limit:        5,
		Offset:       5,
	}
	production, err := testQueries.ListProduction(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, production, 5)

	for _, production := range production {
		require.NotEmpty(t, production)
		require.Equal(t, arg.ProductionID, production.ProductionID)
	}
}
