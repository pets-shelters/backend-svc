CREATE TYPE user_role AS ENUM (
    'manager',
    'employee'
);
CREATE TABLE users (
   id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
   email text NOT NULL,
   shelter_id bigint REFERENCES shelters(id),
   role user_role NOT NULL,
   CONSTRAINT unique_user_email UNIQUE(email)
)