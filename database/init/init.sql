CREATE SCHEMA postgres;

CREATE TABLE postgres.links (
  id VARCHAR PRIMARY KEY,
  type  VARCHAR,
  name  VARCHAR,
  price  VARCHAR
);