CREATE TABLE IF NOT EXISTS todo_staff (
    id VARCHAR(255) NOT NULL,
    todo_id VARCHAR(255) NOT NULL,
    staff_id VARCHAR(255) NOT NULL,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (todo_id) REFERENCES todos(id) ON DELETE CASCADE,
    FOREIGN KEY (staff_id) REFERENCES staffs(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS todo_staff;