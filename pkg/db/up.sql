CREATE TABLE IF NOT EXISTS anywhat (
  id CHAR(27),
  name VARCHAR(24) NOT NULL
  description VARCHAR(1024) DEFAULT NULL,
  createdAt timestamp NULL DEFAULT NULL,
  updatedAt timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY 'id_unique' (id)
);