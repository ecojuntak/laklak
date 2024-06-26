CREATE TABLE applications
(
    id    BIGSERIAL primary key,
    name  varchar(40) NOT NULL,
    team_id BIGINT NOT NULL,
    CONSTRAINT fk_team_id
        FOREIGN KEY(team_id)
            REFERENCES teams(id)
);
