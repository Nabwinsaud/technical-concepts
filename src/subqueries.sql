-- Active: 1689697518705@@127.0.0.1@5432@social

CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name TEXT,
    created_at TIMESTAMP
);

ALTER TABLE customers
ALTER COLUMN created_at
SET DEFAULT CURRENT_TIMESTAMP;

CREATE TABLE IF NOT EXISTS orders (
    order_id SERIAL PRIMARY KEY,
    quantity INT,
    customer_id INT
)

ALTER TABLE orders
ADD FOREIGN KEY (customer_id) REFERENCES customers (id)

INSERT INTO customers (name, age) VALUES ('Matt', 43);

SELECT * FROM customers;

ALTER TABLE customers ADD COLUMN age INT;

-- actions  triggers logs table

CREATE TABLE customers_logs (
    log_id SERIAL PRIMARY KEY,
    customer_id INT,
    action TEXT,
    changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION log_customers_changes()
RETURNS TRIGGER AS $$
BEGIN
INSERT INTO customers_logs(customer_id,action)
VALUES(NEW.id,'INSERT');
RETURN NEW;
END;
$$ LANGUAGE PLPGSQL;

CREATE TRIGGER after_customer_insert
AFTER INSERT on customers FOR EACH ROW
EXECUTE FUNCTION log_customers_changes ();

-- create the virtual tables for customers
CREATE VIEW customer_names AS SELECT name, id, age from customers;

SELECT * FROM customers where created_at > NOW() - INTERVAL '1 sec';

-- some funky concepts

-- let take about concat,collease,cast etc

SELECT
    LTRIM(' Left Trim') as LeftTrim,
    RTRIM('Right trim  ') as RightTrim,
    TRIM('  Both trim ') as AllTrim;

SELECT UPPER('rust') as rustCase, LOWER('RUST') as rustLowerCase;

SELECT LENGTH('rust') as lengthOfRust;

SELECT SUBSTRING('helllo world', 1, 5) as substringHello;

SELECT COALESCE(
        NULL, NULL, NULL, 'i gotcha covered man'
    ) as collease

SELECT CONCAT(
        'Hello', '-', COALESCE(
            NULL, 'no description for strangers'
        )
    ) as description;

CREATE TABLE sample_orders (
    id SERIAL PRIMARY KEY,
    name TEXT,
    description TEXT,
    price DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO
    sample_orders (name, description, price)
VALUES (
        'Item 1',
        'Description 1',
        19.99
    ),
    ('Item 2', NULL, 29.99),
    (
        'Item 3',
        'Description 3',
        NULL
    );

SELECT * from sample_orders;

SELECT price, CONCAT(
        name, ' ', COALESCE(
            description, 'No description available'
        )
    ) as product_details
from sample_orders;

SELECT
    name,
    CASE
        WHEN price IS NULL THEN 'price does not exists'
        -- WHEN price is not null then 'Price is :' || CAST(price as VARCHAR)
        WHEN price is not null then CONCAT(
            'price is :',
            cast(price as VARCHAR)
        )
        ELSE 'no need to check this still fine '
    END as price_status
from sample_orders;