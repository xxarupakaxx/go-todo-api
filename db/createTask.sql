DROP TABLE IF EXISTS task;
CREATE TABLE IF NOT EXISTS task
(
    id
    INT
    UNSIGNED
    NOT
    NULL
    AUTO_INCREMENT,
    created_at
    TIMESTAMP
    NOT
    NULL,
    updated_at
    TIMESTAMP
    NOT
    NULL,
    title
    VARCHAR
(
    255
) NOT NULL ,
    PRIMARY KEY
(
    id
)
    );