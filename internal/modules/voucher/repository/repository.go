package repository

import (
	"database/sql"

	"github.com/yusufekoanggoro/flight-voucher-service/internal/modules/voucher/domain"
)

type VoucherRepository interface {
	FlightExists(flightNumber, date string) (bool, error)
	IsSeatAlreadyUsed(flightNumber, flightDate string, seat string) (bool, error)
	InsertVoucher(v *domain.Voucher) error
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

func (r *voucherRepository) IsSeatAlreadyUsed(flightNumber, flightDate string, seat string) (bool, error) {
	query := `
		SELECT COUNT(*) FROM vouchers 
		WHERE flight_number = ? AND flight_date = ? 
		AND (seat1 = ? OR seat2 = ? OR seat3 = ?)
	`
	var count int
	err := r.db.QueryRow(query, flightNumber, flightDate, seat, seat, seat).Scan(&count)
	return count > 0, err
}

func (r *voucherRepository) InsertVoucher(v *domain.Voucher) error {
	stmt, err := r.db.Prepare(`INSERT INTO vouchers (crew_name, crew_id, flight_number, flight_date, aircraft_type, seat1, seat2, seat3, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		v.CrewName, v.CrewID, v.FlightNumber, v.FlightDate, v.AircraftType,
		v.Seat1, v.Seat2, v.Seat3, v.CreatedAt,
	)
	return err
}
