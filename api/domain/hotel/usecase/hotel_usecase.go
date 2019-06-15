package usecase

import (
	"time"

	"github.com/wincentrtz/bncc/api/domain/hotel"
	"github.com/wincentrtz/bncc/api/models"
)

type hotelUsecase struct {
	hotelRepo      hotel.Repository
	contextTimeout time.Duration
}

func NewHotelUsecase(hr hotel.Repository, timeout time.Duration) hotel.Usecase {
	return &hotelUsecase{
		hotelRepo:      hr,
		contextTimeout: timeout,
	}
}

func (hu *hotelUsecase) FetchHotel() ([]*models.Hotel, error) {
	hotels, err := hu.hotelRepo.FetchHotel()
	if err != nil {
		return nil, err
	}
	return hotels, nil
}
