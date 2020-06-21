CREATE TABLE IF NOT EXISTS onboardings(
  id serial primary key,
  stage integer not null,
  
  role   varchar(17) not null,
  email   varchar(40),
  name varchar(30),
  surname varchar(30)

  address_id integer, -- TODO: ADD STEP
  category_id  integer 
);
