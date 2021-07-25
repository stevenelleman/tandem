-- +migrate Up

CREATE TABLE nodes (
    id varchar(255) UNIQUE NOT NULL,

    edges_out varchar(255)[],
    edge_in varchar(255) NOT NULL,
    scopes varchar(255)[],

    document jsonb NOT NULL,

    -- Metadata
    created_at timestamptz NOT NULL,
    author varchar(255) NOT NULL,
);

CREATE TABLE edges (
    id varchar(255) UNIQUE,

    args jsonb NOT NULL,
    inverse_args jsonb NOT NULL,

    parent_node_id varchar(255),
    child_node_id varchar(255),

    -- Metadata
    created_at timestamptz NOT NULL,
    author varchar(255) NOT NULL,
);

CREATE TABLE scopes (
    id varchar(255) UNIQUE,
    node_ids varchar(255)[],

    -- Metadata
    created_at timestamptz NOT NULL,
    author varchar(255) NOT NULL,
);
