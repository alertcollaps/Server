CREATE TABLE links (
                       short_link text UNIQUE,
                       long_link text,
                       created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


INSERT INTO links (short_link) VALUES
    (1),
    (2);