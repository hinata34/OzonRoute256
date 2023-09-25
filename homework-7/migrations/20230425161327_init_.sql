-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name text NOT NULL DEFAULT '',
    age BIGINT NOT NULL DEFAULT 0
);

CREATE TABLE addresses (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    house_number BIGINT NOT NULL DEFAULT 0,
    street_name text NOT NULL DEFAULT '',
    user_id BIGINT NOT NULL DEFAULT 0
);
--INSERT INTO users(id, name, age) SELECT id, md5(random()::text), random()*(100 - 1) + 1 FROM generate_series(1, 10000) id;
--INSERT INTO addresses(id, house_number, street_name, user_id) SELECT id, random()*(100 - 1) + 1, md5(random()::text), random()*(10000 - 1) + 1 FROM generate_series(1, 10000) id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users, addresses;
-- +goose StatementEnd