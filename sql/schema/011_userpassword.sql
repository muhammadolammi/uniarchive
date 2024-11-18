-- +goose Up
ALTER TABLE users 
ADD COLUMN password  TEXT UNIQUE NOT NULL;


-- +goose Down
ALTER TABLE users 
DROP COLUMN password;