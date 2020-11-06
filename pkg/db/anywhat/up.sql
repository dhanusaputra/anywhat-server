CREATE TABLE IF NOT EXISTS anywhat (
  id SERIAL PRIMARY KEY,
  name VARCHAR(24) NOT NULL,
  description TEXT DEFAULT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);

INSERT INTO anywhat(name, description, created_at, updated_at)
VALUES ('dummyName', 'dummyDesc', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
