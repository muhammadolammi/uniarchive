-- name: CreateDepartment :one

INSERT INTO departments(created_at, updated_at, name, faculty_id)
VALUES(
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetDepartments :many
SELECT * FROM departments;