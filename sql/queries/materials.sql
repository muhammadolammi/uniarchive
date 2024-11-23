-- name: CreateMaterial :one

INSERT INTO materials(created_at, updated_at, name, course_id, cloud_url)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: GetCourseMaterials :many
SELECT * FROM materials WHERE course_id=$1;

-- name: GetDefaultMaterials :many
SELECT materials.id,materials.created_at, materials.updated_at, materials.name , materials.course_id , materials.cloud_url FROM materials
JOIN courses ON courses.id = materials.course_id
 WHERE courses.department_id=$1;