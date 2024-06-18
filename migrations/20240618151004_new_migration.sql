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
-- +goose StatementEnd



-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
