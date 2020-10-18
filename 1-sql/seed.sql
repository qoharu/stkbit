CREATE TABLE IF NOT EXISTS users
(
    id          BIGSERIAL,
    user_name   TEXT,
    parent      BIGINT,
    PRIMARY KEY (id)
);

INSERT INTO users VALUES (DEFAULT, 'Ali', 2);
INSERT INTO users VALUES (DEFAULT, 'Budi', 0);
INSERT INTO users VALUES (DEFAULT, 'Cecep', 1);
