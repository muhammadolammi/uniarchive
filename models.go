package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/muhammadolammi/uniarchive/internal/database"
)

type state struct {
	db        *database.Queries
	PORT      string
	JWTSIGNER string
}

type Course struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `json:"name"`
	CourseCode   string    `json:"course_code"`
	LevelID      uuid.UUID `json:"level_id"`
	DepartmentID uuid.UUID `json:"department_id"`
}

type Department struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	FacultyID uuid.UUID `json:"faculty_id"`
}

type Faculty struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `json:"name"`
	UniversityID uuid.UUID `json:"university_id"`
}

type Level struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Code      int       `json:"code"`
}

type Material struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	CourseID  uuid.UUID `json:"course_id"`
	CloudUrl  string    `json:"cloud_url"`
}

type University struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Alias     string    `json:"alias"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	LevelID      uuid.UUID `json:"level_id"`
	FacultyID    uuid.UUID `json:"faculty_id"`
	DepartmentID uuid.UUID `json:"department_id"`
	// Password     string
	UniversityID uuid.UUID `json:"university_id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	OtherName    string    `json:"other_name"`
	Email        string    `json:"email"`
	MatricNumber string    `json:"matric_number"`
	IsAdmin      bool      `json:"is_admin"`
}
