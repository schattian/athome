CREATE TABLE IF NOT EXISTS identifications(
    id serial primary key,
    user_id integer unique not null,
    dni integer not null,
    verified boolean,
    name varchar(30),
    surname varchar(30),
    license integer,
    tome integer,
    folio integer,
    cue integer
);