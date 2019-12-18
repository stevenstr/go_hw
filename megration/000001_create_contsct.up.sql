CREATE TABLE IF NOT EXISTS contsct(
   id serial PRIMARY KEY NOT NULL,
   first_name varchar(255) NOT NULL,
    last_name varcharg(255) NOT NULL,
    phone varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE
);