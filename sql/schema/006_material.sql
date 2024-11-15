-- +goose Up

CREATE TABLE materials(

 id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,
 name TEXT UNIQUE NOT NULL,
 course_id UUID  NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
 cloud_url TEXT UNIQUE NOT NULL
 
);

-- +goose Down
DROP TABLE materials;