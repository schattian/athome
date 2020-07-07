CREATE TABLE IF NOT EXISTS state_changes(
    id serial primary key,
    stage int not null,
    entity_id int not null,
    entity_table varchar(30) not null,
    created_at timestamp not null,
    name varchar(30) not null
);