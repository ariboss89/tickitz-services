CREATE TABLE users (
  id serial PRIMARY KEY,
  first_name VARCHAR(255),
  last_name VARCHAR(255),
  email VARCHAR(255),
  password TEXT,
  role VARCHAR(255),
  created_at timestamp DEFAULT 'NOW()',
  updated_at timestamp,
  deleted_at timestamp
);