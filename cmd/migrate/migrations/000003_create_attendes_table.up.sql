CREATE TABLE if NOT EXISTS attendes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    event_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) on DELETE CASCADE
    FOREIGN KEY (event_id) REFERENCES events (id) on DELETE CASCADE
);