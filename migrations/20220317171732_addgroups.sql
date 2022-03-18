-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.groups (
    id int primary key GENERATED ALWAYS AS IDENTITY,
    uuid uuid DEFAULT public.uuid_generate_v4(),
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    disabled boolean DEFAULT false,

    name varchar(255) NOT NULL,
    note text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.groups;
-- +goose StatementEnd
