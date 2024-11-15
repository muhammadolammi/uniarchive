-- +goose Up
-- Enable the uuid-ossp extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE  universities(
     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
     created_at TIMESTAMP NOT NULL,
     updated_at TIMESTAMP NOT NULL,
    name TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE universities;
