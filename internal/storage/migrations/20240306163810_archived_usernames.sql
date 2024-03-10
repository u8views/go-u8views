-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_nicknames (
    id           BIGSERIAL NOT NULL PRIMARY KEY,
    user_id      BIGINT    NOT NULL REFERENCES users (id),
    old_nickname VARCHAR   NOT NULL,
    UNIQUE (user_id, old_nickname)
);

CREATE INDEX user_nicknames_idx ON user_nicknames(old_nickname, user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_nicknames;
-- +goose StatementEnd
