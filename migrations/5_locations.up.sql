CREATE TABLE locations (
   id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
   city text NOT NULL,
   address text NOT NULL,
   shelter_id bigint NOT NULL REFERENCES shelters(id)
);