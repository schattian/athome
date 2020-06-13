CREATE TABLE IF NOT EXISTS services(
    id serial primary key,
    user_id int unique not null,
    address_id int,
    calendar_id int not null,
    -- 
    title varchar(30) not null,
    duration_in_minutes int not null,
    price_min int not null,
    price_max int not null
);