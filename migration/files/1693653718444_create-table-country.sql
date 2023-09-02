-- +migrate Up
create table person (
  id serial primary key,
  name varchar,
  country varchar
);

-- +migrate Down
drop table person;
