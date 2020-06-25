CREATE TABLE IF NOT EXISTS purchases(
    id serial primary key,
    user_id int not null,
    src_address_id int,
    distance_in_kilometers float,
    dest_address_id int,
    merchant_id int,
    created_at timestamp not null,
    updated_at timestamp not null,
    items json not null
);

CREATE TABLE IF NOT EXISTS shippings(
    id serial primary key,
    order_price integer not null,
    order_duration_in_minutes integer not null,
    shipping_method_id integer not null,
    event_id integer not null,
    user_id integer not null,
    distance_in_kilometers float,
    real_price integer,
    real_duration_in_minutes integer
);

CREATE TABLE IF NOT EXISTS payments(
    id serial primary key,
    user_id integer not null,
    payment_method_id integer not null,
    card_id integer not null,
    entity_id integer not null,
    entity_table varchar(30) not null,
    amount integer not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    installments integer not null
);

CREATE TABLE IF NOT EXISTS cards(
    id serial primary key,
    user_id integer not null,
    number_hash char(60) not null,
    last_four_digits integer not null,
    cvv_hash char(60) not null,
    expiry_month integer not null,
    expiry_year integer not null,
    holder_dni integer not null,
    holder_name string not null
);

CREATE TABLE IF NOT EXISTS purchase_state_changes(
    id serial primary key,
    stage int not null,
    order_id int not null,
    created_at timestamp not null,
    name varchar(30) not null
);