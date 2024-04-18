CREATE TABLE adopters (
   id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
   name text NOT NULL,
   phone_number char(13) NOT NULL,
   CONSTRAINT unique_adopter_phone UNIQUE(phone_number)
);