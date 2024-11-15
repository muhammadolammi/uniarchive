-- name: CreateUniversity :one

INSERT INTO universities(created_at, updated_at, name, alias)
VALUES(
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUniversities :many
SELECT * FROM universities;

-- name: EditUniversity :exec
UPDATE universities
SET 
  name = COALESCE($1, name),
  alias = COALESCE($2, alias), 
  updated_at=$3
WHERE id = $4;