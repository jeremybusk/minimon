-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.platforms (
    id int primary key GENERATED ALWAYS AS IDENTITY,
    uuid uuid DEFAULT public.uuid_generate_v4(),
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    disabled boolean DEFAULT false,

    name text,
    note text
);

CREATE TABLE public.urls_x_platforms (
    url_id int NOT NULL,
    platform_id int NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.platforms;
DROP TABLE public.urls_x_platforms;
-- +goose StatementEnd
