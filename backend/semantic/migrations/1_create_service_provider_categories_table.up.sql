CREATE TABLE IF NOT EXISTS service_provider_categories(
  id serial primary key,
  name varchar(30) unique not null,
  parent_id uint64
);