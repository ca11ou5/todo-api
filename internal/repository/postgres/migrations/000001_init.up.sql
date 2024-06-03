CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NULL,
    date DATE NOT NULL,
    completed BOOLEAN DEFAULT false
);