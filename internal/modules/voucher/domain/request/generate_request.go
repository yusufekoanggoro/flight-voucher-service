package request

type GenerateRequest struct {
	CrewID       string `json:"crewId" binding:"required"`
	CrewName     string `json:"crewName" binding:"required"`
	FlightNumber string `json:"flightNumber" binding:"required"`
	FlightDate   string `json:"flightDate" binding:"required"`
	AircraftType string `json:"aircraftType" binding:"required"`
}
