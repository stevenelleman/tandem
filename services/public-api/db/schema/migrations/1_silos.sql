-- +migrate Up

CREATE TABLE silos (
    id varchar(255),
    state varchar(255)
);
