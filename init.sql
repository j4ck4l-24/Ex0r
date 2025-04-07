CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    username varchar(255) UNIQUE NOT NULL,
    password_hash varchar(255) NOT NULL,
    email varchar(255) UNIQUE NOT NULL,
    country varchar(10),
    role varchar(255) DEFAULT 'User',
    socials varchar(255),
    website varchar(255),
    access_token varchar(255),
    hidden boolean DEFAULT FALSE,
    banned boolean DEFAULT FALSE,
    team_id int,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (team_id) REFERENCES Teams(id) ON DELETE SET NULL
);

CREATE Table Teams (
    id SERIAL PRIMARY KEY,
    team_name varchar(255) UNIQUE NOT NULL,
    team_password_hash varchar(255) NOT NULL,
    captain_id int UNIQUE,
    hidden boolean DEFAULT FALSE,
    banned boolean DEFAULT FALSE,
    country varchar(10),
    members_id int[],
); 