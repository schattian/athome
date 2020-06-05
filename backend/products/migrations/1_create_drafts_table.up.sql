CREATE TABLE IF NOT EXISTS drafts(
    id serial primary key,
    stage int not null,
    
    title string not null,
    category_id int not null,
);