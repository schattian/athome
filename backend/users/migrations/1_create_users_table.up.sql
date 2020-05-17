CREATE TABLE IF NOT EXISTS users(
  id serial primary key,
  email   varchar(40) not null,
  role   varchar(17) not null,
  name varchar(30) not null,
  surname varchar(30) not null,
  password_hash    char(60) not null
);