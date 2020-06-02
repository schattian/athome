CREATE TABLE IF NOT EXISTS onboarding_identifications(
  id serial primary key,
  onboarding_id integer unique not null,
  dni integer not null,

  name varchar(30),
  surname varchar(30),

  license integer,

  tome integer,
  folio integer,
 
  cue integer,
);
