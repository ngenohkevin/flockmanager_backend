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
WHERE hatchery_id = $1 LIMIT 1;

-- name: ListHatchery :many
SELECT * FROM hatchery
ORDER BY hatchery_id
LIMIT $1
    OFFSET $2;

-- name: UpdateHatchery :one
UPDATE hatchery
SET
    infertile = $1,
    early = $2,
    middle = $3,
    late = $4,
    dead_chicks = $5,
    alive_chicks = $6

WHERE hatchery_id = $1
RETURNING *;