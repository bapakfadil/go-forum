ALTER TABLE users
ADD username VARCHAR(100) NOT NULL;

ALTER TABLE users
ADD CONSTRAINT UNIQUE uniqure_username (username);