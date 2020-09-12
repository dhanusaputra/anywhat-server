CREATE TABLE IF NOT EXISTS anywhat (
  id SERIAL PRIMARY KEY,
  name VARCHAR(24) NOT NULL,
  description TEXT DEFAULT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS user_account (
  id SERIAL PRIMARY KEY,
  username TEXT NOT NULL,
  passwordHash TEXT NOT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL
);

INSERT INTO user_account(username, passwordHash)
VALUES ("admin", "admin");