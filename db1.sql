CREATE DATABASE smarket_db;

CREATE TABLE users (
    id SERIAL PRIMARY KEY ,
    name VARCHAR,
    username VARCHAR,
    password_hash VARCHAR NOT NULL ,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products(
    id SERIAL PRIMARY KEY ,
    name VARCHAR,
    price NUMERIC(10,2),
    quantity INT,
    added_by INT REFERENCES users(id),
    category_id INT REFERENCES category_products(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE category_products(
    id SERIAL PRIMARY KEY ,
    name VARCHAR,
    added_by INT REFERENCES users(id),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP
);

CREATE TABLE sales(
    id SERIAL PRIMARY KEY ,
    user_id INT REFERENCES users(id),
    total_sum NUMERIC(10,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sale_items(
    id SERIAL PRIMARY KEY ,
    sale_id INT REFERENCES sales(id),
    product_id INT REFERENCES products(id),
    quantity INT NOT NULL ,
    price NUMERIC(10,2)
);