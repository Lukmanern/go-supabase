DROP TABLE IF EXISTS todos;
DROP TYPE IF EXISTS status_options;

CREATE TYPE status_options AS ENUM('todo', 'inprogress', 'done');

CREATE TABLE todos (
      id bigint generated always as identity primary key,
      todo text not null,
      status status_options default 'todo',
      created_at timestamp not null default now(),
      deleted_at timestamp default null 
);