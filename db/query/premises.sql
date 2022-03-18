-- name: CreatePremises :one
INSERT INTO premises (
    farm,
    house
) VALUES (
             $1, $2
         ) RETURNING *;

-- name: GetPremises :one
SELECT * FROM premises
WHERE premises_id = $1 LIMIT 1;

-- name: ListPremises :many
SELECT * FROM premises
ORDER BY premises_id
LIMIT $1
OFFSET $2;
