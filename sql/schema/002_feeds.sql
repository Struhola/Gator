-- +goose Up
CREATE TABLE feeds (
id UUID primary key,
created_at TIMESTAMP not null,
updated_at TIMESTAMP not null,
last_fetched_at TIMESTAMP,
name Text not null,
url Text unique not null,
user_id UUID 
    REFERENCES users(id)
    ON DELETE CASCADE
    NOT NULL
);
-- +goose Down
DROP TABLE feeds; 