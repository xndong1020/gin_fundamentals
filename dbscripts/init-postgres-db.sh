#!/bin/bash
set -e

echo "creating db"
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE docker;
    GRANT ALL PRIVILEGES ON DATABASE docker TO root;

    \c docker

    CREATE SCHEMA hollywood;

    SET SCHEMA 'hollywood';

    CREATE TABLE IF NOT EXISTS hollywood.albums (
        id serial NOT NULL,
        title VARCHAR(255),
        artist VARCHAR(255),
        price DECIMAL,
        CONSTRAINT "PK_tbl_albums" PRIMARY KEY (id)
    );

    INSERT INTO hollywood.albums ("title", "artist", "price") VALUES
    ('Blue Train','John Coltrane',56.99)
    ,('Jeru','Gerry Mulligan',17.99)
    ,('Sarah Vaughan and Clifford Brown','Sarah Vaughan',39.99)
    ;
EOSQL