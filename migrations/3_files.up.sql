CREATE TABLE files (
   id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
   bucket text NOT NULL,
   path text NOT NULL,
   CONSTRAINT unique_file UNIQUE(bucket, path)
)

CREATE TABLE temporary_files (
   id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
   file_id bigint NOT NULL REFERENCES files(id),
   user_id bigint NOT NULL REFERENCES users(id),
   created_at timestamp without time zone NOT NULL,
   CONSTRAINT unique_temporary_file UNIQUE(file_id)
)