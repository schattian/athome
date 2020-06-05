CREATE TABLE IF NOT EXISTS drafts(
    id serial primary key,
    user_id int unique not null,
    stage int not null
);