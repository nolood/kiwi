-- +goose Up
-- +goose StatementBegin

CREATE TYPE session AS ENUM ('fill_profile_age', 'fill_profile_photo', 'fill_profile_gender', 'fill_profile_about', 'fill_blacklist', 'none');

CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	telegram_id BIGINT NOT NULL UNIQUE,
	first_name VARCHAR(255),
	last_name VARCHAR(255),
	username VARCHAR(255) NOT NULL UNIQUE,
	language_code VARCHAR(5),
	is_premium BOOLEAN NOT NULL DEFAULT FALSE,
	photo_url VARCHAR(1000),
	session session NOT NULL DEFAULT 'none' 
);


CREATE TABLE IF NOT EXISTS profiles (
	id SERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL UNIQUE,
	user_tg_id BIGINT NOT NULL UNIQUE,
	age INTEGER,
	gender VARCHAR(255),
	photo_id VARCHAR(2000),
	about VARCHAR(255),
	is_active BOOLEAN NOT NULL DEFAULT FALSE,
	CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
	CONSTRAINT fk_user_tg FOREIGN KEY (user_tg_id) REFERENCES users(telegram_id)
);
-- +goose StatementEnd



-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS profiles CASCADE;
DROP TYPE IF EXISTS session;
-- +goose StatementEnd
