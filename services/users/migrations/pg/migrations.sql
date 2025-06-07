CREATE SCHEMA users;

-- Table: users
CREATE TABLE users.users (
  id UUID PRIMARY KEY,
  email VARCHAR UNIQUE,
  password_hash VARCHAR,
  name VARCHAR,
  role VARCHAR,
  is_active BIT,
  created_by VARCHAR,
  updated_by VARCHAR,
  deleted_by VARCHAR,
  deleted_at TIMESTAMP NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);