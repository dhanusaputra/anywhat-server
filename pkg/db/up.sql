CREATE TABLE IF NOT EXISTS anywhat (
  id SERIAL PRIMARY KEY,
  name VARCHAR(24) NOT NULL,
  description TEXT DEFAULT NULL,
  createdAt timestamp NULL DEFAULT NULL,
  updatedAt timestamp NULL DEFAULT NULL
);