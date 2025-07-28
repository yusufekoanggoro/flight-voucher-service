package request

type CheckFlightRequest struct {
	FlightNumber string `json:"flightNumber" binding:"required"`
	Date         string `json:"date" binding:"required"` // Format YYYY-MM-DD
}
