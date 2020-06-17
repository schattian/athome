CREATE TABLE IF NOT EXISTS notifications(
    id serial primary key,
    user_id int not null,
    priority int not null,
    body text not null,
    entity_id int,
    entity_table int,
    created_at timestamp not null,
    received_at timestamp,
    seen_at timestamp
);
