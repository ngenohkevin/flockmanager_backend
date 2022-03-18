// Code generated by sqlc. DO NOT EDIT.
// source: rainbowrooster.sql

package db

import (
	"context"
)

const createRainbowRooster = `-- name: CreateRainbowRooster :one
INSERT INTO rainbowrooster (
    title
) VALUES (
          $1
) RETURNING id, title, created_at
`

func (q *Queries) CreateRainbowRooster(ctx context.Context, title string) (Rainbowrooster, error) {
	row := q.queryRow(ctx, q.createRainbowRoosterStmt, createRainbowRooster, title)
	var i Rainbowrooster
	err := row.Scan(&i.ID, &i.Title, &i.CreatedAt)
	return i, err
}

const deleteRainbowRooster = `-- name: DeleteRainbowRooster :exec
DELETE FROM rainbowrooster
WHERE id = $1
`

func (q *Queries) DeleteRainbowRooster(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteRainbowRoosterStmt, deleteRainbowRooster, id)
	return err
}

const getRainbowRooster = `-- name: GetRainbowRooster :one
SELECT id, title, created_at FROM rainbowrooster
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRainbowRooster(ctx context.Context, id int32) (Rainbowrooster, error) {
	row := q.queryRow(ctx, q.getRainbowRoosterStmt, getRainbowRooster, id)
	var i Rainbowrooster
	err := row.Scan(&i.ID, &i.Title, &i.CreatedAt)
	return i, err
}

const listRainbowRooster = `-- name: ListRainbowRooster :many
SELECT id, title, created_at FROM rainbowrooster
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListRainbowRoosterParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListRainbowRooster(ctx context.Context, arg ListRainbowRoosterParams) ([]Rainbowrooster, error) {
	rows, err := q.query(ctx, q.listRainbowRoosterStmt, listRainbowRooster, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Rainbowrooster
	for rows.Next() {
		var i Rainbowrooster
		if err := rows.Scan(&i.ID, &i.Title, &i.CreatedAt); err != nil {
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
