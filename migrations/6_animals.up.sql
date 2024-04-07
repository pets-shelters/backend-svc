CREATE TYPE animal_gender AS ENUM (
    'female',
    'male'
);

CREATE TYPE animal_type AS ENUM (
    'кіт',
    'пес'
);

CREATE OR REPLACE FUNCTION add_animal_type(new_value text) RETURNS VOID AS $$
BEGIN
    EXECUTE 'ALTER TYPE animal_type ADD VALUE IF NOT EXISTS ''' || new_value || '''';
END;
$$ LANGUAGE plpgsql;

CREATE TABLE animals (
   id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
   location_id bigint NOT NULL REFERENCES locations(id),
   photo bigint NOT NULL REFERENCES files(id),
   name text NOT NULL,
   birth_date date NOT NULL,
   type animal_type NOT NULL,
   gender animal_gender NOT NULL,
   sterilized boolean NOT NULL,
   private_description text,
   public_description text
);