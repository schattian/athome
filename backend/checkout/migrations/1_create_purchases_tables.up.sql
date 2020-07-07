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