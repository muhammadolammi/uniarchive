-- +goose Up

CREATE TABLE users(

 id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,
 first_name TEXT  NOT NULL,
 last_name TEXT NOT NULL,
 other_name TEXT NOT NULL,
 email TEXT UNIQUE NOT NULL,
 matric_number TEXT UNIQUE NOT NULL,
 level_id UUID  NOT NULL,
  FOREIGN KEY (level_id)   REFERENCES levels(id) ,
 faculty_id UUID  NOT NULL,
  FOREIGN KEY (faculty_id)   REFERENCES faculties(id) , 
 department_id UUID NOT NULL,
  FOREIGN KEY (department_id)   REFERENCES departments(id),
 university_id UUID NOT NULL,
  FOREIGN KEY (university_id)   REFERENCES universities(id)
 );

-- +goose Down
DROP TABLE users;