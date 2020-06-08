CREATE TABLE IF NOT EXISTS draft_lines(
    id serial primary key,
    draft_id int not null,
    title string not null,
    category_id int not null,
    price int,
    stock int,
    image_ids varchar(50) []
);