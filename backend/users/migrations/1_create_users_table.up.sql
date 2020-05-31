CREATE TABLE IF NOT EXISTS users(
  id serial primary key,
  email   varchar(40) not null,
  category   varchar(40),
  role   varchar(17) not null,
  name varchar(30) not null,
  surname varchar(30) not null,
  password_hash    char(60) not null
);

ALTER TABLE ONLY users ADD CONSTRAINT uq_users UNIQUE(email, role);