CREATE TABLE shelters (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name text NOT NULL,
    logo text NOT NULL,
    city text NOT NULL,
    phone_number char(12) NOT NULL,
    instagram text,
    facebook text,
    created_at date NOT NULL
);

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