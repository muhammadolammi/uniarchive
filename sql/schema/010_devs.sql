-- +goose Up

CREATE TABLE devs(

 id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,
 name TEXT UNIQUE NOT NULL,
 password TEXT UNIQUE NOT NULL
 );

-- +goose Down
DROP TABLE devs;