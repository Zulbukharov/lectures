CREATE TABLE IF NOT EXISTS posts (
  id            SERIAL,
  author_id     bigint NOT NULL,
  content       text NOT NULL
);
