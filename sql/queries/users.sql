-- name: CreateUser :one

INSERT INTO users(created_at, updated_at, first_name, last_name, other_name, email, matric_number, password, level_id,  faculty_id, department_id, university_id)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12
)
RETURNING *;

-- name: GetUsers :many
SELECT * FROM users;

-- name: MakeUserAnAdmin :exec
UPDATE users 
SET is_admin=true WHERE id=$1;

-- name: GetUserWithEmail :one
SELECT * FROM users WHERE email=$1;

-- name: GetUserWithMatricNumber :one
SELECT * FROM users WHERE matric_number=$1;