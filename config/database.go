package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func init() {

	fmt.Println("Conectando a:", "bitacora.db")

	var err error
	Database, err = sql.Open("sqlite3", "bitacora.db")
	if err != nil {
		log.Fatal("Error creando conexión: ", err)
	}

	// Validar conexión real
	if err := Database.Ping(); err != nil {
		log.Fatal("Error conectando a SQLite: ", err)
	}

	createTables := `
	CREATE TABLE IF NOT EXISTS records (
		id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		name TEXT NOT NULL,
		lab TEXT NOT NULL,
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
`

	_, err = Database.Exec(createTables)
	if err != nil {
		log.Fatal("Error creando tabla principal: ", err)
	}

	log.Println("✅ Base de datos conectada con éxito y tablas listas")
}
