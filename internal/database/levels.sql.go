// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: levels.sql

package database

import (
	"context"
	"time"
)

const createLevel = `-- name: CreateLevel :one

INSERT INTO levels(created_at, updated_at, name, code)
VALUES(
    $1,
    $2,
    $3,
    $4
)
RETURNING id, created_at, updated_at, name, code
`

type CreateLevelParams struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Code      int32
}

func (q *Queries) CreateLevel(ctx context.Context, arg CreateLevelParams) (Level, error) {
	row := q.db.QueryRowContext(ctx, createLevel,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Code,
	)
	var i Level
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Code,
	)
	return i, err
}

const getLevels = `-- name: GetLevels :many
SELECT id, created_at, updated_at, name, code FROM levels
`

func (q *Queries) GetLevels(ctx context.Context) ([]Level, error) {
	rows, err := q.db.QueryContext(ctx, getLevels)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Level
	for rows.Next() {
		var i Level
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Code,
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