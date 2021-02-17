CREATE DATABASE fapidb;
\c fapidb;

CREATE TABLE IF NOT EXISTS users(
    uid bigserial PRIMARY KEY,
    nickname varchar(15) NOT NULL UNIQUE,
    email varchar(40) NOT NULL UNIQUE,
    password varchar(100 NOT NULL,
    status char(1) default '0',
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS wallets(
    public_key varchar(32) PRIMARY KEY /*UNIQUE NOT NULL*/,
    usr bigint NOT NULL,
    balance real DEFAULT 0.0,
    updated_at timestamp DEFAULT current_timestamp,
    constraint wallets_usr_fk FOREIGN KEY(usr)
        REFERENCES users(uid)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);