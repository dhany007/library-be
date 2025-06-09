-- Enum for transaction status
CREATE TYPE books.transaction_status AS ENUM (
  'borrowed',
  'returned',
  'cancelled'
);

-- Table: transactions
CREATE TABLE books.transactions (
  id UUID PRIMARY KEY,
  user_id UUID REFERENCES users(id),
  book_id UUID REFERENCES books(id),
  borrowed_at TIMESTAMP,
  due_date TIMESTAMP,
  returned_at TIMESTAMP NULL,
  status transaction_status,
  is_active BOOLEAN,
  created_by VARCHAR,
  updated_by VARCHAR,
  deleted_by VARCHAR,
  deleted_at TIMESTAMP NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);