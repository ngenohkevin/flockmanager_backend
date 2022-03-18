-- name: CreateRainbowRooster :one
INSERT INTO rainbowrooster (
    title
) VALUES (
          $1
) RETURNING *;

-- name: GetRainbowRooster :one
SELECT * FROM rainbowrooster
WHERE id = $1 LIMIT 1;

-- name: ListRainbowRooster :many
SELECT * FROM rainbowrooster
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteRainbowRooster :exec
DELETE FROM rainbowrooster
WHERE id = $1;