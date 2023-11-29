CREATE TABLE IF NOT EXISTS link (
                                    full_link VARCHAR(255) NOT NULL UNIQUE,
                                    short_link VARCHAR(10) NOT NULL UNIQUE PRIMARY KEY
);