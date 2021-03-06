// Code generated by sqlc. DO NOT EDIT.
// source: hatchery.sql

package db

import (
	"context"
	"database/sql"
)

const createHatchery = `-- name: CreateHatchery :one
INSERT INTO hatchery (
    hatchery_id,
    infertile,
    early,
    middle,
    late,
    dead_chicks,
    alive_chicks
) VALUES (
             $1, $2, $3, $4, $5, $6, $7
         ) RETURNING id, hatchery_id, created_at, infertile, early, middle, late, dead_chicks, alive_chicks
`

type CreateHatcheryParams struct {
	HatcheryID  int64         `json:"hatcheryID"`
	Infertile   sql.NullInt64 `json:"infertile"`
	Early       sql.NullInt64 `json:"early"`
	Middle      sql.NullInt64 `json:"middle"`
	Late        sql.NullInt64 `json:"late"`
	DeadChicks  sql.NullInt64 `json:"deadChicks"`
	AliveChicks sql.NullInt64 `json:"aliveChicks"`
}

func (q *Queries) CreateHatchery(ctx context.Context, arg CreateHatcheryParams) (Hatchery, error) {
	row := q.queryRow(ctx, q.createHatcheryStmt, createHatchery,
		arg.HatcheryID,
		arg.Infertile,
		arg.Early,
		arg.Middle,
		arg.Late,
		arg.DeadChicks,
		arg.AliveChicks,
	)
	var i Hatchery
	err := row.Scan(
		&i.ID,
		&i.HatcheryID,
		&i.CreatedAt,
		&i.Infertile,
		&i.Early,
		&i.Middle,
		&i.Late,
		&i.DeadChicks,
		&i.AliveChicks,
	)
	return i, err
}

const getHatchery = `-- name: GetHatchery :one
SELECT id, hatchery_id, created_at, infertile, early, middle, late, dead_chicks, alive_chicks FROM hatchery
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetHatchery(ctx context.Context, id int64) (Hatchery, error) {
	row := q.queryRow(ctx, q.getHatcheryStmt, getHatchery, id)
	var i Hatchery
	err := row.Scan(
		&i.ID,
		&i.HatcheryID,
		&i.CreatedAt,
		&i.Infertile,
		&i.Early,
		&i.Middle,
		&i.Late,
		&i.DeadChicks,
		&i.AliveChicks,
	)
	return i, err
}

const listHatchery = `-- name: ListHatchery :many
SELECT id, hatchery_id, created_at, infertile, early, middle, late, dead_chicks, alive_chicks FROM hatchery
WHERE hatchery_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListHatcheryParams struct {
	HatcheryID int64 `json:"hatcheryID"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *Queries) ListHatchery(ctx context.Context, arg ListHatcheryParams) ([]Hatchery, error) {
	rows, err := q.query(ctx, q.listHatcheryStmt, listHatchery, arg.HatcheryID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Hatchery
	for rows.Next() {
		var i Hatchery
		if err := rows.Scan(
			&i.ID,
			&i.HatcheryID,
			&i.CreatedAt,
			&i.Infertile,
			&i.Early,
			&i.Middle,
			&i.Late,
			&i.DeadChicks,
			&i.AliveChicks,
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
