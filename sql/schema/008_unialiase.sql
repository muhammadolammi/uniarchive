-- +goose Up
ALTER TABLE universities 
ADD COLUMN alias TEXT NOT NULL DEFAULT 'DEFAULT-ALIAS';


-- +goose Down
ALTER TABLE universities 
DROP COLUMN alias;