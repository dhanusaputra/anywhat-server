CREATE TABLE IF NOT EXISTS anywhat (
  id SERIAL PRIMARY KEY,
  name VARCHAR(24) NOT NULL,
  description TEXT DEFAULT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL
);