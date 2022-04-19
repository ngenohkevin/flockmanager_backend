-- name: CreateKuroiler :one
INSERT INTO kuroiler (
    title,
    house
) VALUES (
          $1, $2
) RETURNING *;

-- name: GetKuroiler :one
SELECT * FROM kuroiler
WHERE id = $1 LIMIT 1;

-- name: ListKuroiler :many
SELECT * FROM kuroiler
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateKuroiler :one
UPDATE kuroiler
SET house = $2
WHERE id = $1
RETURNING *;

-- name: DeleteKuroiler :exec
DELETE FROM kuroiler
WHERE id = $1;