-- init.sql

CREATE TABLE IF NOT EXISTS data_tank (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    `time` timestamp,
    Barel INTEGER NOT NULL,
	Volume INTEGER NOT NULL,
    Temperature REAL NOT NULL,
    WaterDebit REAL NOT NULL,
    Alarm varchar(5000) NOT NULL
);
