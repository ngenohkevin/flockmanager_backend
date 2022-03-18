// Code generated by sqlc. DO NOT EDIT.
// source: hatchery.sql

package db

import (
	"context"
	"database/sql"
)

const createHatchery = `-- name: CreateHatchery :one
INSERT INTO hatchery (
        infertile,
        early,
        middle,
        late,
        dead_chicks,
        alive_chicks
) VALUES (
             $1, $2, $3, $4, $5, $6
         ) RETURNING hatchery_id, created_at, infertile, early, middle, late, dead_chicks, alive_chicks, kuroiler_id, rainbowrooster_id, broiler_id, layers_id
`

type CreateHatcheryParams struct {
	Infertile   sql.NullInt32 `json:"infertile"`
	Early       sql.NullInt32 `json:"early"`
	Middle      sql.NullInt32 `json:"middle"`
	Late        sql.NullInt32 `json:"late"`
	DeadChicks  sql.NullInt32 `json:"deadChicks"`
	AliveChicks sql.NullInt64 `json:"aliveChicks"`
}

func (q *Queries) CreateHatchery(ctx context.Context, arg CreateHatcheryParams) (Hatchery, error) {
	row := q.queryRow(ctx, q.createHatcheryStmt, createHatchery,
		arg.Infertile,
		arg.Early,
		arg.Middle,
		arg.Late,
		arg.DeadChicks,
		arg.AliveChicks,
	)
	var i Hatchery
	err := row.Scan(
		&i.HatcheryID,
		&i.CreatedAt,
		&i.Infertile,
		&i.Early,
		&i.Middle,
		&i.Late,
		&i.DeadChicks,
		&i.AliveChicks,
		&i.KuroilerID,
		&i.RainbowroosterID,
		&i.BroilerID,
		&i.LayersID,
	)
	return i, err
}

const getHatchery = `-- name: GetHatchery :one
SELECT hatchery_id, created_at, infertile, early, middle, late, dead_chicks, alive_chicks, kuroiler_id, rainbowrooster_id, broiler_id, layers_id FROM hatchery
WHERE hatchery_id = $1 LIMIT 1
`

func (q *Queries) GetHatchery(ctx context.Context, hatcheryID sql.NullInt64) (Hatchery, error) {
	row := q.queryRow(ctx, q.getHatcheryStmt, getHatchery, hatcheryID)
	var i Hatchery
	err := row.Scan(
		&i.HatcheryID,
		&i.CreatedAt,
		&i.Infertile,
		&i.Early,
		&i.Middle,
		&i.Late,
		&i.DeadChicks,
		&i.AliveChicks,
		&i.KuroilerID,
		&i.RainbowroosterID,
		&i.BroilerID,
		&i.LayersID,
	)
	return i, err
}

const listHatchery = `-- name: ListHatchery :many
SELECT hatchery_id, created_at, infertile, early, middle, late, dead_chicks, alive_chicks, kuroiler_id, rainbowrooster_id, broiler_id, layers_id FROM hatchery
ORDER BY hatchery_id
LIMIT $1
    OFFSET $2
`

type ListHatcheryParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListHatchery(ctx context.Context, arg ListHatcheryParams) ([]Hatchery, error) {
	rows, err := q.query(ctx, q.listHatcheryStmt, listHatchery, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Hatchery
	for rows.Next() {
		var i Hatchery
		if err := rows.Scan(
			&i.HatcheryID,
			&i.CreatedAt,
			&i.Infertile,
			&i.Early,
			&i.Middle,
			&i.Late,
			&i.DeadChicks,
			&i.AliveChicks,
			&i.KuroilerID,
			&i.RainbowroosterID,
			&i.BroilerID,
			&i.LayersID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateHatchery = `-- name: UpdateHatchery :one
UPDATE hatchery
SET
    infertile = $1,
    early = $2,
    middle = $3,
    late = $4,
    dead_chicks = $5,
    alive_chicks = $6

WHERE hatchery_id = $1
RETURNING hatchery_id, created_at, infertile, early, middle, late, dead_chicks, alive_chicks, kuroiler_id, rainbowrooster_id, broiler_id, layers_id
`

type UpdateHatcheryParams struct {
	Infertile   sql.NullInt32 `json:"infertile"`
	Early       sql.NullInt32 `json:"early"`
	Middle      sql.NullInt32 `json:"middle"`
	Late        sql.NullInt32 `json:"late"`
	DeadChicks  sql.NullInt32 `json:"deadChicks"`
	AliveChicks sql.NullInt64 `json:"aliveChicks"`
}

func (q *Queries) UpdateHatchery(ctx context.Context, arg UpdateHatcheryParams) (Hatchery, error) {
	row := q.queryRow(ctx, q.updateHatcheryStmt, updateHatchery,
		arg.Infertile,
		arg.Early,
		arg.Middle,
		arg.Late,
		arg.DeadChicks,
		arg.AliveChicks,
	)
	var i Hatchery
	err := row.Scan(
		&i.HatcheryID,
		&i.CreatedAt,
		&i.Infertile,
		&i.Early,
		&i.Middle,
		&i.Late,
		&i.DeadChicks,
		&i.AliveChicks,
		&i.KuroilerID,
		&i.RainbowroosterID,
		&i.BroilerID,
		&i.LayersID,
	)
	return i, err
}
