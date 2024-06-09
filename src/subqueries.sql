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