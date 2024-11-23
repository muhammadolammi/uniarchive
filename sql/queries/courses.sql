
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

-- name: GetUserCourses :many
SELECT courses.id, courses.created_at, courses.updated_at, courses.name, courses.course_code, courses.level_id, courses.department_id
FROM courses
JOIN users ON users.department_id = courses.department_id AND users.level_id = courses.level_id
WHERE users.id = sqlc.arg('user_id');
