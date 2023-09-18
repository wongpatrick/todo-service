CREATE SCHEMA todoapp ;

CREATE TABLE todoapp.user (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL
);


CREATE TABLE todoapp.task (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES todoapp.user(id) ON DELETE CASCADE,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  status VARCHAR(50) NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  modified_at TIMESTAMP DEFAULT NOW()
);