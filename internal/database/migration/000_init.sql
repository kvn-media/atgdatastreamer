-- init.sql

CREATE TABLE IF NOT EXISTS data_tanks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    capacity INTEGER NOT NULL,
    temperature REAL
);
