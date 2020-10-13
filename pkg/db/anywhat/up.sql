CREATE TABLE IF NOT EXISTS anywhat (
  id SERIAL PRIMARY KEY,
  name VARCHAR(24) NOT NULL,
  description TEXT DEFAULT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL
);

INSERT INTO anywhat(name, description, created_at, updated_at)
VALUES ('dummyName', 'dummyDesc', '2020-08-05 18:19:03', '2020-08-05 18:19:03');
