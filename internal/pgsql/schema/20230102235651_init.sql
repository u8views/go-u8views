-- +goose Up
-- +goose StatementBegin
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
    time    TIMESTAMP NOT NULL,
    user_id BIGINT    NOT NULL REFERENCES users (id),
    count   BIGINT    NOT NULL,
    PRIMARY KEY (time, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE profile_hourly_views_stats;
DROP TABLE profile_total_views;
DROP TABLE users;
-- +goose StatementEnd
