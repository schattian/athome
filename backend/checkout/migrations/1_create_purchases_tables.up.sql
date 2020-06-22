CREATE TABLE IF NOT EXISTS purchases(
    id serial primary key,
    user_id int not null,
    src_address_id int,
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
    event_id integer not null,

    real_price integer,
    real_duration_in_minutes integer 
)

CREATE TABLE IF NOT EXISTS purchase_state_changes(
    id serial primary key,
    stage int not null,
    order_id int not null,
    created_at timestamp not null,
    name varchar(30) not null
);