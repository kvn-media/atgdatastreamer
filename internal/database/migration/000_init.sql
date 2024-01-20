-- init.sql

CREATE TABLE IF NOT EXISTS data_tank (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    time timestamp,
    Barel BIGINT NOT NULL,
	VolumeBarel INTEGER NOT NULL,
    AveTemperature INTEGER NOT NULL,
    WaterDebit DECIMAL(10, 2) NOT NULL,
	TempProduct INTEGER NOT NULL,
    Alarm varchar(5000) NOT NULL
);
