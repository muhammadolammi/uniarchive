-- +goose Up

CREATE TABLE levels(

 id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,
 name TEXT UNIQUE NOT NULL, -- eg "100 level"
 code INT UNIQUE NOT NULL  -- eg 100

);

-- +goose Down
DROP TABLE levels;