-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS username_history
(
    id                 BIGSERIAL                NOT NULL PRIMARY KEY,
    user_id            BIGINT                   NOT NULL REFERENCES users (id),
    social_provider    SOCIAL_PROVIDER          NOT NULL,
    canonical_username VARCHAR                  NOT NULL,
    created_at         TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at         TIMESTAMP WITH TIME ZONE NOT NULL,
    UNIQUE (canonical_username, social_provider, user_id)
);

INSERT INTO username_history (user_id, social_provider, canonical_username, created_at, updated_at)
SELECT id, social_provider, canonical_username, created_at, last_login_at
FROM users
ON CONFLICT (canonical_username, social_provider, user_id) DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE username_history;
-- +goose StatementEnd
