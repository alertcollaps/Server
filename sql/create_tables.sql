CREATE TABLE quote (
                       id integer UNIQUE,
                       cash integer CHECK (cash >= 0) DEFAULT 0,
                       created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);

INSERT INTO quote (id) VALUES
    (1),
    (2);