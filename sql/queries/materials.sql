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

-- name: GetMaterials :many
SELECT * FROM materials;