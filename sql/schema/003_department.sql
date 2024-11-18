-- +goose Up
CREATE TABLE departments(
     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT UNIQUE NOT NULL,
    faculty_id UUID NOT NULL ,
    FOREIGN KEY (faculty_id)  REFERENCES faculties(id) ON DELETE CASCADE

);

-- +goose Down
DROP TABLE departments;