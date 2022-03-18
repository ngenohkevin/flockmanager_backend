-- name: CreateProduction :one
INSERT INTO production (
    eggs,
    dirty,
    wrong_shape,
    weak_shell,
    damaged,
    hatching_eggs
) VALUES (
             $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetProduction :one
SELECT * FROM production
WHERE production_id = $1 LIMIT 1;

-- name: ListProduction :many
SELECT * FROM production
ORDER BY production_id
LIMIT $1
OFFSET $2;

-- name: UpdateProduction :one
UPDATE production
SET
     eggs = $1,
     dirty = $2,
     wrong_shape = $3,
     weak_shell = $4,
     damaged = $5,
     hatching_eggs = $6

WHERE production_id = $1
RETURNING *;
