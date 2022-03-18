package db

import (
	"context"
	"github.com/ngenohkevin/flockmanager_backend/db/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomProduction(t *testing.T) Production {
	args := CreateProductionParams{
		Eggs:         int64(util.RandomEggs()),
		Dirty:        util.RandomDirty(),
		WrongShape:   util.RandomWrongShape(),
		WeakShell:    util.RandomWeakShell(),
		Damaged:      util.RandomDamaged(),
		HatchingEggs: int64(util.RandomHatchingEggs()),
	}

	prod, err := testQueries.CreateProduction(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, prod)

	require.Equal(t, args.Eggs, prod.Eggs)
	require.Equal(t, args.Dirty, prod.Dirty)
	require.Equal(t, args.WrongShape, prod.WrongShape)
	require.Equal(t, args.WeakShell, prod.WeakShell)
	require.Equal(t, args.Damaged, prod.Damaged)
	require.Equal(t, args.HatchingEggs, prod.HatchingEggs)

	require.NotZero(t, prod.ProductionID)
	require.NotZero(t, prod.CreatedAt)

	return prod
}

func TestCreateProduction(t *testing.T) {
	CreateRandomProduction(t)
}
func TestGetProduction(t *testing.T) {
	prod1 := CreateRandomProduction(t)

	prod2, err := testQueries.GetProduction(context.Background(), prod1.ProductionID)
	require.NoError(t, err)
	require.NotEmpty(t, prod2)

	require.Equal(t, prod1.ProductionID, prod2.ProductionID)
	require.Equal(t, prod1.Eggs, prod2.Eggs)
	require.Equal(t, prod1.Dirty, prod2.Dirty)
	require.Equal(t, prod1.WrongShape, prod2.WrongShape)
	require.Equal(t, prod1.WeakShell, prod2.WeakShell)
	require.Equal(t, prod1.Damaged, prod2.Damaged)
	require.Equal(t, prod1.HatchingEggs, prod2.HatchingEggs)
	require.WithinDuration(t, prod1.CreatedAt, prod2.CreatedAt, time.Second)
}

func TestListProduction(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomProduction(t)
	}
	arg := ListProductionParams{
		Limit:  5,
		Offset: 5,
	}
	prod, err := testQueries.ListProduction(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, prod, 5)

	for _, prod := range prod {
		require.NotEmpty(t, prod)
	}
}
