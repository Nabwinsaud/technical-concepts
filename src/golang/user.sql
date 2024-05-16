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