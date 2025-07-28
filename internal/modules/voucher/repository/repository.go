package repository

import (
	"database/sql"
	"time"
)

type VoucherRepository interface {
	FlightExists(flightNumber, date string) (bool, error)
	InsertVoucher(name, id, flightNumber, date, aircraft string, seats []string) error
}

type voucherRepository struct {
	db *sql.DB
}

func NewVoucherRepository(db *sql.DB) VoucherRepository {
	return &voucherRepository{db}
}

func (r *voucherRepository) FlightExists(flightNumber string, date string) (bool, error) {
	query := `
		SELECT COUNT(*) FROM vouchers 
		WHERE flight_number = ? AND flight_date = ?
	`

	var count int
	err := r.db.QueryRow(query, flightNumber, date).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *voucherRepository) InsertVoucher(name, id, flightNumber, date, aircraft string, seats []string) error {
	stmt, err := r.db.Prepare(`INSERT INTO vouchers (crew_name, crew_id, flight_number, flight_date, aircraft_type, seat1, seat2, seat3, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name, id, flightNumber, date, aircraft, seats[0], seats[1], seats[2], time.Now().Format(time.RFC3339))
	return err
}
