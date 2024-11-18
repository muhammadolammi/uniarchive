-- name: CreateUser :one

INSERT INTO users(created_at, updated_at, name,level_id,  faculty_id, department_id, university_id)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *;

-- name: GetUsers :many
SELECT * FROM users;

-- name: MakeUserAnAdmin :exec
UPDATE users 
SET is_admin=true WHERE id=$1;