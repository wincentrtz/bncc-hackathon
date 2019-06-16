package usecase

import (
	"time"

	"github.com/wincentrtz/bncc/api/domain/flight"
	"github.com/wincentrtz/bncc/api/models"
)

type flightUsecase struct {
	flightRepo     flight.Repository
	contextTimeout time.Duration
}

func NewFlightUsecase(fr flight.Repository, timeout time.Duration) flight.Usecase {
	return &flightUsecase{
		flightRepo:     fr,
		contextTimeout: timeout,
	}
}

func (hu *flightUsecase) FetchFlight() ([]*models.Flight, error) {
	flights, err := hu.flightRepo.FetchFlight()
	if err != nil {
		return nil, err
	}
	return flights, nil
}
