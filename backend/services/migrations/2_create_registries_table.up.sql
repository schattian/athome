CREATE TABLE IF NOT EXISTS registries(
    id serial primary key,
    user_id int unique not null,
    stage int not null,
    --
    address_id int,
    --
    title varchar(30),
    duration_in_minutes int,
    price_min int,
    price_max int,
    --
    calendar_id int
);
