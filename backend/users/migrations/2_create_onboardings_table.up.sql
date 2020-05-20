CREATE TABLE IF NOT EXISTS onboardings(
  id serial primary key,
  email   varchar(40),
  role   varchar(17) not null,
  stage integer not null,
  name varchar(30),
  surname varchar(30)
);
