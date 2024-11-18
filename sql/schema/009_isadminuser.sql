-- +goose Up
ALTER TABLE universities 
ADD COLUMN is_admin  BOOLEAN NOT NULL DEFAULT false;


-- +goose Down
ALTER TABLE universities 
DROP COLUMN is_admin;