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