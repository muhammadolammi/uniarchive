-- name: CreateLevel :one

INSERT INTO levels(created_at, updated_at, name, code)
VALUES(
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetLevels :many
SELECT * FROM levels;