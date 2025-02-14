CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  uuid VARCHAR(255) NOT NULL,
  username VARCHAR(20) NOT NULL,
  email VARCHAR(100),
  ip_address VARCHAR(15) NOT NULL,
  country VARCHAR(50) NOT NULL,
  createdAt VARCHAR(255) NOT NULL,

  UNIQUE (uuid),
  UNIQUE (email),
  UNIQUE (username)
);
