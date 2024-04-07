CREATE TABLE temporary_files (
     id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
     file_id bigint NOT NULL REFERENCES files(id) ON DELETE CASCADE,
     user_id bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
     created_at timestamp without time zone NOT NULL,
     CONSTRAINT unique_temporary_file UNIQUE(file_id)
);