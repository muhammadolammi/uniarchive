-- +goose Up
-- Enable the uuid-ossp extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE  faculties(
id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,
 name TEXT UNIQUE NOT NULL,
 university_id UUID NOT NULL ,
 FOREIGN KEY (university_id) REFERENCES universities(id) ON DELETE CASCADE

);

-- +goose Down
DROP TABLE  faculties;