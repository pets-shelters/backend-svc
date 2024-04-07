CREATE TABLE files (
   id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
   bucket text NOT NULL,
   path text NOT NULL,
   CONSTRAINT unique_file UNIQUE(bucket, path)
);