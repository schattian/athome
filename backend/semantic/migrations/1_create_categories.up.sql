CREATE TABLE IF NOT EXISTS service_provider_categories(
    id serial primary key,
    name varchar(30) unique not null,
    parent_id int
);

CREATE TABLE IF NOT EXISTS merchant_categories(
    id serial primary key,
    name varchar(30) unique not null,
    parent_id int
);

CREATE TABLE IF NOT EXISTS product_categories(
    id serial primary key,
    name varchar(30) unique not null,
    parent_id int
);