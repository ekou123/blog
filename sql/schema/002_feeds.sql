-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(30) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    url VARCHAR(100) NOT NULL,
    user_id UUID NOT NULL,
    UNIQUE(url),
    CONSTRAINT users_fk
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;