-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS timescaledb;

CREATE TABLE users
(
    id BIGSERIAL NOT NULL PRIMARY KEY
);

CREATE TABLE profile_total_views
(
    user_id BIGINT NOT NULL REFERENCES users (id) PRIMARY KEY,
    count   BIGINT NOT NULL
);

CREATE TABLE profile_hourly_views_stats
(
    user_id BIGINT    NOT NULL REFERENCES users (id),
    time    TIMESTAMP NOT NULL,
    count   BIGINT    NOT NULL,
    PRIMARY KEY (user_id, time)
);

SELECT create_hypertable('profile_hourly_views_stats', 'time');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE profile_hourly_views_stats;
DROP TABLE profile_total_views;
DROP TABLE users;
-- +goose StatementEnd
