-- name: CreateKuroiler :one
INSERT INTO kuroiler (
    title
) VALUES (
          $1
) RETURNING *;

-- name: GetKuroiler :one
SELECT * FROM kuroiler
WHERE id = $1 LIMIT 1;

-- name: ListKuroiler :many
SELECT * FROM kuroiler
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteKuroiler :exec
DELETE FROM kuroiler
WHERE id = $1;