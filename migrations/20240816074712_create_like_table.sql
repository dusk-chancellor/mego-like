-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS likes
(
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    post_id INTEGER NULL,
    comment_id INTEGER NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS likes;
-- +goose StatementEnd
