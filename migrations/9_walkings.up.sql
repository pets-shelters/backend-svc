CREATE TYPE walking_status AS ENUM (
    'pending',
    'approved'
);

CREATE TABLE walkings (
   id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
   status walking_status NOT NULL,
   animal_id bigint NOT NULL REFERENCES animals(id),
   name text NOT NULL,
   phone_number char(13) NOT NULL,
   date date NOT NULL,
   time time,
   CONSTRAINT time_required_for_approved CHECK ((walkings.status = 'pending') OR (walkings.time IS NOT NULL))
);