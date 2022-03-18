-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.http_connections (
    id int primary key GENERATED ALWAYS AS IDENTITY,
    uuid uuid DEFAULT public.uuid_generate_v4(),
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    disabled boolean DEFAULT false,

    url_id int,
    domain_name_lookup_time bigint,
    tcp_connection_time bigint,
    connect_time bigint,
    pre_transfer_time bigint,
    start_transfer_time bigint,
    server_processing_time bigint,
    tls_handshake_time bigint,
    context_transfer_time bigint,
    total_time bigint,
    start_time timestamp with time zone,
    stop_time timestamp with time zone,
    status_code bigint,
    ip_address inet,
    response_body_regex_match boolean,
    response_header_regex_match boolean,
    response_body_text text,
    response_header_text text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.http_connections;
-- +goose StatementEnd
