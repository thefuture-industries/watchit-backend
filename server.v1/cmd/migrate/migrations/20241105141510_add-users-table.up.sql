CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  uuid VARCHAR(255) NOT NULL,
  secret_word VARCHAR(255) NOT NULL,
  username VARCHAR(100) NOT NULL,
  username_upper VARCHAR(100) NOT NULL,
  email VARCHAR(50),
  email_upper VARCHAR(50),
  ip_address VARCHAR(40) NOT NULL,
  country VARCHAR(70) NOT NULL,
  regionName VARCHAR(70) NOT NULL,
  zip VARCHAR(40) NOT NULL,
  createdAt VARCHAR(100) NOT NULL,

  UNIQUE (email),
  UNIQUE (uuid),
  UNIQUE (secret_word)
);
