-- +migrate Up

CREATE TABLE silos (
    id varchar(255) UNIQUE,
    state varchar(255)
);
