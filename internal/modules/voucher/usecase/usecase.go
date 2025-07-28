package usecase

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/yusufekoanggoro/flight-voucher-service/internal/modules/voucher/domain/request"
	"github.com/yusufekoanggoro/flight-voucher-service/internal/modules/voucher/domain/response"
	"github.com/yusufekoanggoro/flight-voucher-service/internal/modules/voucher/repository"
)

type VoucherUsecase interface {
	CheckFlight(req request.CheckFlightRequest) (bool, error)
	GenerateVoucher(req request.GenerateRequest) (response.GenerateVoucherResponse, error)
}

type voucherUsecase struct {
	repo repository.VoucherRepository
}

func NewVoucherUsecase(repo repository.VoucherRepository) VoucherUsecase {
	uc := &voucherUsecase{
		repo: repo,
	}

	return uc
}

func (u *voucherUsecase) CheckFlight(req request.CheckFlightRequest) (bool, error) {
	return u.repo.FlightExists(req.FlightNumber, req.FlightDate)
}

func (u *voucherUsecase) GenerateVoucher(req request.GenerateRequest) (response.GenerateVoucherResponse, error) {
	availableSeats, err := getAvailableSeats(req.Aircraft)
	if err != nil {
		return response.GenerateVoucherResponse{}, err
	}

	selectedSeats := randomSeats(availableSeats, 3)

	for _, seat := range selectedSeats {
		exists, err := u.repo.IsSeatAlreadyUsed(req.FlightNumber, req.FlightDate, seat)
		if err != nil {
			return response.GenerateVoucherResponse{}, err
		}
		if exists {
			return response.GenerateVoucherResponse{}, fmt.Errorf("seat %s already used on flight %s at %s", seat, req.FlightNumber, req.FlightDate)
		}
	}

	err = u.repo.InsertVoucher(req.Name, req.ID, req.FlightNumber, req.FlightDate, req.Aircraft, selectedSeats)
	if err != nil {
		return response.GenerateVoucherResponse{}, err
	}

	return response.GenerateVoucherResponse{
		Seats: selectedSeats,
	}, nil
}

func getAvailableSeats(aircraft string) ([]string, error) {
	seatMap := map[string][]string{
		"ATR":            generateSeats(1, 18, []string{"A", "C", "D", "F"}),
		"Airbus 320":     generateSeats(1, 32, []string{"A", "B", "C", "D", "E", "F"}),
		"Boeing 737 Max": generateSeats(1, 32, []string{"A", "B", "C", "D", "E", "F"}),
	}

	seats, ok := seatMap[aircraft]
	if !ok {
		return nil, fmt.Errorf("unsupported aircraft type: %s", aircraft)
	}

	return seats, nil
}

func generateSeats(start, end int, letters []string) []string {
	var seats []string
	for i := start; i <= end; i++ {
		for _, l := range letters {
			seats = append(seats, fmt.Sprintf("%d%s", i, l))
		}
	}

	return seats
}

func randomSeats(all []string, count int) []string {
	if len(all) < count {
		return nil // atau error
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	shuffled := make([]string, len(all))
	copy(shuffled, all)
	r.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled[:count]
}
