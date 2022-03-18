-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.http_connection_trigger (
    id int primary key GENERATED ALWAYS AS IDENTITY,
    uuid uuid DEFAULT public.uuid_generate_v4(),
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    disabled boolean DEFAULT false,
    group_id int DEFAULT 1,

    name varchar(255) NOT NULL,
    domain_name_lookup_time time DEFAULT '00:00:00.5',
    tcp_connection_time time DEFAULT '00:00:00.5',
    connect_time time DEFAULT '00:00:00.5',
    pretransfer_time time DEFAULT '00:00:00.5',
    start_transfer_time time DEFAULT '00:00:00.5',
    server_processing_time time DEFAULT '00:00:00.5',
    tls_handshake_time time DEFAULT '00:00:00.5',
    context_transfer_time time DEFAULT '00:00:00.5',
    total_time time DEFAULT '00:00:00.5',
    status_code int,
    check_response_body boolean DEFAULT false,
    check_response_header boolean DEFAULT false,
    response_body_regex text,
    response_header_regex text,
    ip_address inet DEFAULT '0.0.0.0/0'::inet,
    note text,
    FOREIGN KEY(group_id) REFERENCES groups(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.http_connection_trigger;
-- +goose StatementEnd
