CREATE TABLE teams
(
    id    BIGSERIAL primary key,
    name  varchar(40) NOT NULL UNIQUE
);
