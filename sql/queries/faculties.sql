-- name: CreateFaculty :one

INSERT INTO faculties(created_at, updated_at, name, university_id)
VALUES(
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetFaculties :many
SELECT * FROM faculties;