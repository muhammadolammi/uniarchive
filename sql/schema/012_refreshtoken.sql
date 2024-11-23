-- +goose Up
ALTER TABLE users 
ADD COLUMN refresh_token  TEXT UNIQUE ;


-- +goose Down
ALTER TABLE users 
DROP COLUMN refresh_token;