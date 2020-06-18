CREATE TABLE IF NOT EXISTS calendars(
    id serial primary key,
    name varchar(30) not null,
    user_id int not null,
    group_id int
);

CREATE TABLE IF NOT EXISTS availabilities(
    id serial primary key,
    day_of_week integer not null,
    start_hour integer not null,
    start_minute integer not null,
    end_hour integer not null,
    end_minute integer not null
);

CREATE TABLE IF NOT EXISTS events(
    id serial primarty key,
    claimant_id integer not null,
    calendar_id integer not null,
    order_id integer unique not null,
    is_confirmed boolean not null,
    day_of_week integer not null,
    start_hour integer not null,
    start_minute integer not null,
    end_hour integer not null,
    end_minute integern ot null
);