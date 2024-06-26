CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users
(
    id            uuid DEFAULT uuid_generate_v4() primary key,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);
