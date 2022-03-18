-- +goose Up
-- +goose StatementBegin
INSERT INTO urls (path) VALUES ('https://example.com');
INSERT INTO urls (path) VALUES ('https://example.org');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
