-- name: CreateHatchery :one
INSERT INTO hatchery (
    infertile,
    early,
    middle,
    late,
    dead_chicks,
    alive_chicks
) VALUES (
             $1, $2, $3, $4, $5, $6
         ) RETURNING *;

-- name: GetHatchery :one
SELECT * FROM hatchery
WHERE id = $1 LIMIT 1;

-- name: ListHatchery :many
SELECT * FROM hatchery
WHERE hatchery_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;