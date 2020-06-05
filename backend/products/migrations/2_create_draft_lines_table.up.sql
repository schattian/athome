CREATE TABLE IF NOT EXISTS draft_lines(
    id serial primary key,
    title string not null,
    category_id int not null
);