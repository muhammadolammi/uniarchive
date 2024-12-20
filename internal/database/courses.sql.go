// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: courses.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createCourse = `-- name: CreateCourse :one

INSERT INTO courses(created_at, updated_at, name,course_code,level_id, department_id)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING id, created_at, updated_at, name, course_code, level_id, department_id
`

type CreateCourseParams struct {
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Name         string
	CourseCode   string
	LevelID      uuid.UUID
	DepartmentID uuid.UUID
}

func (q *Queries) CreateCourse(ctx context.Context, arg CreateCourseParams) (Course, error) {
	row := q.db.QueryRowContext(ctx, createCourse,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.CourseCode,
		arg.LevelID,
		arg.DepartmentID,
	)
	var i Course
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.CourseCode,
		&i.LevelID,
		&i.DepartmentID,
	)
	return i, err
}

const getUserCourses = `-- name: GetUserCourses :many
SELECT courses.id, courses.created_at, courses.updated_at, courses.name, courses.course_code, courses.level_id, courses.department_id
FROM courses
JOIN users ON users.department_id = courses.department_id AND users.level_id = courses.level_id
WHERE users.id = $1
`

func (q *Queries) GetUserCourses(ctx context.Context, userID uuid.UUID) ([]Course, error) {
	rows, err := q.db.QueryContext(ctx, getUserCourses, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Course
	for rows.Next() {
		var i Course
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.CourseCode,
			&i.LevelID,
			&i.DepartmentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
