
-- +migrate Up
INSERT INTO users(name,email) VALUES('webonise','weboniselab@info.com');
INSERT INTO users(name,email) VALUES('test','test@info.com');
-- +migrate Down
DELETE FROM users WHERE email in('weboniselab@info.com','test@info.com');
