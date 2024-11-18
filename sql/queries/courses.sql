
-- name: CreateCourse :one

INSERT INTO courses(created_at, updated_at, name,course_code,level_id, department_id)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;
-- name: GetCourses :many
SELECT * FROM courses;