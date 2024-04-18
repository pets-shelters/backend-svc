CREATE TABLE shelters (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name text NOT NULL,
    logo bigint NOT NULL REFERENCES files(id),
    phone_number char(13) NOT NULL,
    instagram text,
    facebook text,
    created_at date NOT NULL
);