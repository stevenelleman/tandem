-- +migrate Up

CREATE TABLE nodes (
    nodeId varchar(255) UNIQUE,
    doc varchar(255),
    edgesOut varchar(255)[],
    edgesIn varchar(255)[],
    date varchar(255),
    author varchar(255),
);

CREATE TABLE edges (
    edgeId varchar(255) UNIQUE,
    sourceDoc varchar(255),
    targetDoc varchar(255),
    args varchar(255),
    inverseArgs varchar(255),
    sourceNodeId varchar(255),
    targetNodeId varchar(255),
    date varchar(255),
    author varchar(255)
);

CREATE TABLE scopes (
    scopeId varchar(255) UNIQUE,
    nodes varchar(255)[],
    edges varchar(255)[]
);
