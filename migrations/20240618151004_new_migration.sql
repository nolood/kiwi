-- +goose Up
-- +goose StatementBegin

CREATE TYPE session AS ENUM ('fill_profile_age', 'fill_profile_name', 'fill_profile_location', 'fill_profile_photo', 'fill_profile_gender', 'fill_profile_about', 'fill_blacklist', 'none');

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


CREATE TABLE IF NOT EXISTS cities (
                                      geonameid INT PRIMARY KEY,
                                      name VARCHAR(200),
                                      asciiname VARCHAR(200),
                                      alternatenames TEXT,
                                      latitude DOUBLE PRECISION,
                                      longitude DOUBLE PRECISION,
                                      feature_class CHAR(1),
                                      feature_code VARCHAR(10),
                                      country_code CHAR(2),
                                      cc2 VARCHAR(60),
                                      admin1_code VARCHAR(20),
                                      admin2_code VARCHAR(80),
                                      admin3_code VARCHAR(20),
                                      admin4_code VARCHAR(20),
                                      population INT,
                                      elevation INT,
                                      dem INT,
                                      timezone VARCHAR(40),
                                      modification_date DATE
);

-- +goose StatementEnd



-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS profiles CASCADE;
DROP TYPE IF EXISTS session;
DROP TABLE IF EXISTS cities;
-- +goose StatementEnd
