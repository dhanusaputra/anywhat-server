CREATE TABLE IF NOT EXISTS user_account (
  id SERIAL PRIMARY KEY,
  username TEXT UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  last_login_at timestamp
);

INSERT INTO user_account(username, password_hash, created_at, updated_at)
VALUES ('admin', '$2y$10$5lCjHTMJqRvjZ.jFlADGK.1iLh78xEJerrUuFavSE5HsGfYMPZ8cG', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
