-- +goose Up
CREATE TABLE department(
     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT UNIQUE NOT NULL,
    faculty_id UUID NOT NULL  REFERENCES faculty(id) ON DELETE CASCADE

);

-- +goose Down
DROP TABLE department;