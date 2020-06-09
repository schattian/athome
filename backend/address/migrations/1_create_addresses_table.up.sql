CREATE TABLE IF NOT EXISTS addresses(
    id serial primary key,
    user_id int not null,
    country varchar(30) not null,
    province varchar(30) not null,
    number integer not null,
    street varchar(30) not null,
    floor integer,
    department integer,
    latitude decimal(9,6) not null,
    longitude decimal(9,6) not null,
    alias varchar(30) 
);

ALTER TABLE ONLY addresses ADD CONSTRAINT uq_addresses UNIQUE(latitude, longitude, user_id);

