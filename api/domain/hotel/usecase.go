package hotel

import (
	"github.com/wincentrtz/bncc/api/models"
)

type Usecase interface {
	FetchHotel() ([]*models.Hotel, error)
}
