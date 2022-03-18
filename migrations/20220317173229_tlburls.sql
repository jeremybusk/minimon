-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.urls (
    id int primary key GENERATED ALWAYS AS IDENTITY,
    uuid uuid DEFAULT public.uuid_generate_v4(),
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    disabled boolean DEFAULT false,
    group_id int DEFAULT 1,
    http_connection_trigger_id int DEFAULT 1,

    note text,
    path varchar(2048) UNIQUE NOT NULL,
    allow_insecure_tls boolean,
    FOREIGN KEY(group_id) REFERENCES groups(id),
    FOREIGN KEY(http_connection_trigger_id) REFERENCES http_connection_trigger(id)
);


CREATE TABLE public.urls_x_groups (
    url_id int NOT NULL,
    group_id int NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.urls;
DROP TABLE public.urls_x_groups;
-- +goose StatementEnd
