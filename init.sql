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
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
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
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    members_id int[],
);

CREATE TABLE Challenges (
    id SERIAL PRIMARY KEY,
    chall_name VARCHAR(255) NOT NULL,
    chall_desc TEXT,
    category VARCHAR(255) NOT NULL,
    current_points INT,
	initial_points INT,
    min_points INT,
    max_attempts INT,
    type VARCHAR(255),
    hidden BOOLEAN DEFAULT TRUE,
    author_name VARCHAR(255),
    decay_type VARCHAR(255),
    decay_value INT,
    connection_string TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    requirements VARCHAR(256),
    next_chall_id INT,
    FOREIGN KEY (next_chall_id) REFERENCES Challenges(id) ON DELETE SET NULL
);