-- +migrate Up

CREATE TABLE nodes (
    id varchar(255) UNIQUE,
    doc varchar(255)
);

CREATE TABLE edges (
    id varchar(255) UNIQUE,
    txn_id varchar(255),
    source_doc varchar(255),
    args varchar(255),
    date varchar(255),
    author varchar(255)
);

CREATE TABLE scopes (
    id varchar(255) UNIQUE,
    nodes varchar(255)[],
    edges varchar(255)[]
);
