-- name: CreateBroiler :one
INSERT INTO broiler (
    title
) VALUES (
          $1
) RETURNING *;

-- name: GetBroiler :one
SELECT * FROM broiler
WHERE id = $1 LIMIT 1;

-- name: ListBroiler :many
SELECT * FROM broiler
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteBroiler :exec
DELETE FROM broiler
WHERE id = $1;