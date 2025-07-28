package usecase

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/yusufekoanggoro/voucher-seat-be/internal/modules/voucher/domain/request"
	"github.com/yusufekoanggoro/voucher-seat-be/internal/modules/voucher/domain/response"
	"github.com/yusufekoanggoro/voucher-seat-be/internal/modules/voucher/repository"
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
	return u.repo.FlightExists(req.FlightNumber, req.Date)
}

func (u *voucherUsecase) GenerateVoucher(req request.GenerateRequest) (response.GenerateVoucherResponse, error) {
	availableSeats, err := getAvailableSeats(req.Aircraft)
	if err != nil {
		return response.GenerateVoucherResponse{}, err
	}

	rand.Seed(time.Now().UnixNano())
	selectedSeats := randomSeats(availableSeats, 3)

	err = u.repo.InsertVoucher(req.Name, req.ID, req.FlightNumber, req.Date, req.Aircraft, selectedSeats)
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
	r := make([]string, 0, count)
	m := make(map[string]bool)
	for len(r) < count {
		s := all[rand.Intn(len(all))]
		if !m[s] {
			r = append(r, s)
			m[s] = true
		}
	}
	return r
}
