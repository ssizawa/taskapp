CREATE TABLE IF NOT EXISTS User(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(32) UNIQUE NOT NULL,
    password VARCHAR(64) NOT NULL DEFAULT 'project443'
);

INSERT IGNORE INTO USER(name) VALUES
    ('konishi kyosuke'),
    ('izawa shin');