package request

type CheckFlightRequest struct {
	FlightNumber string `json:"flightNumber" binding:"required"`
	FlightDate   string `json:"flightDate" binding:"required"` // Format YYYY-MM-DD
}
