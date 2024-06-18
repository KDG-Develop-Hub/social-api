CREATE TABLE IF NOT EXISTS users
(
    id           INT PRIMARY KEY,
    username     VARCHAR(64) NOT NULL,
    display_name VARCHAR(64) NOT NULL,
    password     VARCHAR(64) NOT NULL,
    email        VARCHAR(64) NOT NULL,
    avatar       VARCHAR(64),
    bio          TEXT,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);