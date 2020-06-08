CREATE TABLE IF NOT EXISTS products(
    id serial primary key,
    user_id int not null,
    title string not null,
    category_id int not null,
    price int not null,
    stock int not null,
    image_ids varchar(50) []
);