
CREATE SCHEMA books

-- Table: authors
CREATE TABLE books.authors (
  id UUID PRIMARY KEY,
  name VARCHAR,
  bio TEXT,
  is_active BIT,
  created_by VARCHAR,
  updated_by VARCHAR,
  deleted_by VARCHAR,
  deleted_at TIMESTAMP NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

-- Table: categories
CREATE TABLE books.categories (
  id UUID PRIMARY KEY,
  name VARCHAR,
  description TEXT,
  is_active BIT,
  created_by VARCHAR,
  updated_by VARCHAR,
  deleted_by VARCHAR,
  deleted_at TIMESTAMP NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

-- Table: books
CREATE TABLE books.books (
  id UUID PRIMARY KEY,
  title VARCHAR,
  isbn VARCHAR,
  stock INT,
  author_id UUID REFERENCES books.authors(id),
  category_id UUID REFERENCES books.categories(id),
  description TEXT,
  is_active BIT,
  created_by VARCHAR,
  updated_by VARCHAR,
  deleted_by VARCHAR,
  deleted_at TIMESTAMP NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);


CREATE INDEX idx_books_isbn ON books.books(isbn);
