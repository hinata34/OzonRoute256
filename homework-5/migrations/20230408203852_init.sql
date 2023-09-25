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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users, addresses;
-- +goose StatementEnd
