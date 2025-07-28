package infrastucture

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitDB(path string) *sql.DB {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Fatal("failed to open db: ", err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS vouchers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		crew_name TEXT,
		crew_id TEXT,
		flight_number TEXT,
		flight_date TEXT,
		aircraft_type TEXT,
		seat1 TEXT,
		seat2 TEXT,
		seat3 TEXT,
		created_at TEXT,
		UNIQUE(flight_number, flight_date, seat1),
		UNIQUE(flight_number, flight_date, seat2),
		UNIQUE(flight_number, flight_date, seat3)
	);`

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal("failed to create table: ", err)
	}

	return db
}
