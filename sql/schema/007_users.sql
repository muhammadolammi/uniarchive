-- +goose Up

CREATE TABLE users(

 id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,
 name TEXT UNIQUE NOT NULL,
 level_id UUID  NOT NULL REFERENCES levels(id) ,
 faculty_id UUID  NOT NULL REFERENCES faculty(id) , 
 department_id UUID NOT NULL REFERENCES department(id),
 university_id UUID NOT NULL REFERENCES universities(id)
 );

-- +goose Down
DROP TABLE users;