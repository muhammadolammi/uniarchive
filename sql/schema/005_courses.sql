-- +goose Up

CREATE TABLE courses(

 id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,
 name TEXT UNIQUE NOT NULL,
 course_code TEXT UNIQUE NOT NULL,
 level_id UUID  NOT NULL,
  FOREIGN KEY (level_id)   REFERENCES levels(id) ,
 department_id UUID NOT NULL ,
  FOREIGN KEY (department_id)   REFERENCES departments(id) ON DELETE  CASCADE
);

-- +goose Down
DROP TABLE courses;