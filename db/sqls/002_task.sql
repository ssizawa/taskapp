CREATE TABLE IF NOT EXISTS Task(
    id INT AUTO_INCREMENT PRIMARY KEY,
    task_name VARCHAR(128) NOT NULL UNIQUE,
    description TEXT,
    pic VARCHAR(32) NOT NULL,
    FOREIGN KEY fk_pic(pic) REFERENCES User(name) ON DELETE CASCADE ON UPDATE CASCADE,
    deadline DATE,
    status VARCHAR(16) NOT NULL DEFAULT 'todo'
);