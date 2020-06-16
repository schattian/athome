CREATE TABLE IF NOT EXISTS messages(
    id serial primary key,
    receiver_id int not null,
    sender_id int not null,
    body text not null,
    created_at timestamp not null,
    received_at timestamp,
    seen_at timestamp
);