CREATE TABLE tasks (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  description text NOT NULL,
  start_date date NOT NULL,
  end_date date NOT NULL,
  time time,
  CONSTRAINT end_date_later CHECK (end_date >= start_date)
);

CREATE TABLE tasks_animals (
   id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
   animal_id bigint NOT NULL REFERENCES animals(id) ON DELETE CASCADE,
   task_id bigint NOT NULL REFERENCES tasks(id) ON DELETE CASCADE
);

CREATE TABLE tasks_executions (
   id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
   task_id bigint NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
   user_id bigint REFERENCES users(id) ON DELETE SET NULL,
   date date NOT NULL,
   done_at timestamp without time zone NOT NULL
);