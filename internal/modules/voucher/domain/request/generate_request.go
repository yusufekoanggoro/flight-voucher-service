package request

type GenerateRequest struct {
	ID           string `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	FlightNumber string `json:"flightNumber" binding:"required"`
	FlightDate   string `json:"date" binding:"required"`
	Aircraft     string `json:"aircraft" binding:"required"`
}
