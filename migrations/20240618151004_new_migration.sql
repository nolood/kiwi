-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	telegram_id BIGINT NOT NULL UNIQUE,
	first_name VARCHAR(255),
	last_name VARCHAR(255),
	username VARCHAR(255) NOT NULL UNIQUE,
	language_code VARCHAR(5),
	is_premium BOOLEAN NOT NULL DEFAULT FALSE,
	photo_url VARCHAR(1000)
);

CREATE TABLE IF NOT EXISTS profiles (
	id SERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL UNIQUE,
	age INTEGER,
	gender VARCHAR(255),
	about VARCHAR(255),
	CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +goose StatementEnd



-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS profiles;
-- +goose StatementEnd
