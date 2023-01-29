-- +goose Up
-- +goose StatementBegin
CREATE TYPE SOCIAL_PROVIDER AS ENUM (
    'github',
    'gitlab',
    'bitbucket'
    );

CREATE TABLE users
(
    id                      BIGSERIAL                NOT NULL PRIMARY KEY,
    social_provider         SOCIAL_PROVIDER          NOT NULL,
    social_provider_user_id VARCHAR                  NOT NULL,
    username                VARCHAR                  NOT NULL,
    name                    VARCHAR                  NOT NULL,
    canonical_username      VARCHAR                  NOT NULL,
    created_at              TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at              TIMESTAMP WITH TIME ZONE NOT NULL,
    last_login_at           TIMESTAMP WITH TIME ZONE NOT NULL,
    UNIQUE (social_provider, social_provider_user_id)
);

-- Duplicates allowed because username can be changed
CREATE INDEX SOCIAL_PROVIDER_USERNAME ON users (social_provider, canonical_username);

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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE profile_hourly_views_stats;
DROP TABLE profile_total_views;
DROP TABLE users;

DROP TYPE SOCIAL_PROVIDER;
-- +goose StatementEnd
