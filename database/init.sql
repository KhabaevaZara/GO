CREATE TABLE IF NOT EXISTS Users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS Tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT DEFAULT 'todo',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(id)
);

INSERT INTO Users (username, email) VALUES 
('test_user', 'user@example.com');

-- Заполнение 100 задач
WITH RECURSIVE generate_series(value) AS (
    SELECT 1
    UNION ALL
    SELECT value+1 FROM generate_series WHERE value+1<=100
)
INSERT INTO Tasks (user_id, title, description)
SELECT 1, 'Task ' || value, 'Description for task ' || value
FROM generate_series;