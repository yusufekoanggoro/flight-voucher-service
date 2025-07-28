package domain

type Voucher struct {
	ID           int64
	CrewName     string
	CrewID       string
	FlightNumber string
	FlightDate   string // YYYY-MM-DD
	AircraftType string
	Seat1        string
	Seat2        string
	Seat3        string
	CreatedAt    string
}
