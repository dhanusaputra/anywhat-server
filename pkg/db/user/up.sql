CREATE TABLE IF NOT EXISTS user_account (
  id SERIAL PRIMARY KEY,
  username TEXT NOT NULL,
  password_hash TEXT NOT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  last_login_at timestamp NULL DEFAULT NULL
);

INSERT INTO user_account(username, password_hash)
VALUES ('admin', '$2y$10$5lCjHTMJqRvjZ.jFlADGK.1iLh78xEJerrUuFavSE5HsGfYMPZ8cG');
