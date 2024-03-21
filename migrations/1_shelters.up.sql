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