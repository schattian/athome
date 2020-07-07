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