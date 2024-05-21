-- Active: 1714124661583@@127.0.0.1@5432@godb@public

CREATE TABLE IF Not exists users (
    id serial PRIMARY key,
    name Text,
    password TEXT,
    roles TEXT []
)

select * from users

ALTER TABLE users ADD COLUMN age int

INSERT INTO
    users (name, password, age, roles)
VALUES (
        'RAM',
        'Admin@123',
        21,
        ARRAY['admin', 'read-only']
    )
    -- teams

CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(200),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)

CREATE TABLE IF NOT EXISTS channel (
    id SERIAL PRIMARY key,
    channel_name TEXT,
    post_id int,
    Foreign Key (post_id) REFERENCES posts (id)
)

INSERT INTO
    posts (name)
VALUES (
        'Marketing is the ultimate iceberg'
    )

SELECT *
FROM posts po
    join channel c on po.id = c.post_id
order by created_at ASC

INSERT INTO
    channel (channel_name, post_id)
VALUES ('Marketing', 2)

ALTER TABLE users ADD CONSTRAINT unique_name UNIQUE (name);