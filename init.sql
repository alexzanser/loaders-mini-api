CREATE TABLE IF NOT EXISTS customers (
    id                  SERIAL  PRIMARY KEY,
    username            varchar NOT NULL UNIQUE,
    passwd_hash         varchar NOT NULL,
    balance             integer
);

CREATE TABLE IF NOT EXISTS loaders (
    id                  SERIAL  PRIMARY KEY,
    username            varchar NOT NULL UNIQUE,
    passwd_hash         varchar NOT NULL,
    max_weight          integer,
    drunk               boolean,
    fatigue             integer,
    salary              integer,
    balance             integer,
    completed_tasks     integer[]
);

CREATE TABLE IF NOT EXISTS tasks (
    id                  SERIAL  PRIMARY KEY,
    customer_id         integer REFERENCES customers ON DELETE CASCADE,
    name                varchar NOT NULL,
    weight              integer,
    completed           boolean DEFAULT false
);
