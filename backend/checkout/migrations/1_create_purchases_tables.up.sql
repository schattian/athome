CREATE TABLE IF NOT EXISTS purchases(
    id serial primary key,
    user_id int not null,
    address_id int,
    merchant_id int,
    created_at timestamp not null,
    updated_at timestamp not null,
    items json not null
);

CREATE TABLE IF NOT EXISTS purchase_state_changes(
    id serial primary key,
    stage int not null,
    order_id int not null,
    created_at timestamp not null,
    name varchar(30) not null
);