-- +goose Up
-- +goose StatementBegin
INSERT INTO groups (name) VALUES ('Default');
INSERT INTO http_connection_trigger (name) VALUES ('Default');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
