-- +goose Up
-- +goose StatementBegin
CREATE TABLE referrals
(
    referee_user_id  BIGINT NOT NULL REFERENCES users (id) PRIMARY KEY,
    referrer_user_id BIGINT NOT NULL REFERENCES users (id)
);

CREATE INDEX referrals_referee_user_id_idx ON referrals (referee_user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE referrals;
-- +goose StatementEnd
