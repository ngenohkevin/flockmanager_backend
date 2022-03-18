-- name: CreateLayers :one
INSERT INTO layers (
    title
) VALUES (
          $1
) RETURNING *;

-- name: GetLayers :one
SELECT * FROM layers
WHERE id = $1 LIMIT 1;

-- name: ListLayers :many
SELECT * FROM layers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteLayers :exec
DELETE FROM layers
WHERE id = $1;