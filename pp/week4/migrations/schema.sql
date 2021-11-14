CREATE TABLE IF NOT EXISTS orders (
  id UUID PRIMARY KEY,
  SKU text NOT NULL,
  price int NOT NULL
);
CREATE TABLE IF NOT EXISTS files (
  id SERIAL PRIMARY KEY,
  file bytea NOT NULL,
  file_name text NOT NULL
);