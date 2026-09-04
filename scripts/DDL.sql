-- This SQL script file includes all the table creation functions - triggers - function - etc

CREATE TABLE IF NOT EXISTS records (
		id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		name TEXT NOT NULL,
		lab TEXT NOT NULL,
		endUser VARCHAR NOT NULL,
		equipment TEXT NOT NULL,
		startDateTime INTEGER NOT NULL,
		endDateTime INTEGER,
		received INTEGER,
		returned INTEGER,
		comments TEXT,
		timestamp INTEGER,
		formated_startDateTime TEXT,
		formated_endDateTime TEXT,
		formated_timestamp TEXT
);


CREATE TRIGGER IF NOT EXISTS format_dates_after_insert
	AFTER INSERT ON records
	FOR EACH ROW
	BEGIN
		UPDATE records
		SET
			formated_startDateTime = datetime(NEW.startDateTime / 1000, 'unixepoch', 'localtime'),
			formated_endDateTime = CASE
				WHEN NEW.endDateTime IS NOT NULL AND NEW.endDateTime != 0 THEN
					datetime(NEW.endDateTime / 1000, 'unixepoch', 'localtime')
				ELSE NULL
			END,
			formated_timestamp = datetime(NEW.timestamp, 'unixepoch', 'localtime')
		WHERE id = NEW.id;
END;