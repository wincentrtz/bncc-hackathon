package hotel

import (
	"github.com/wincentrtz/bncc/api/models"
)

type Repository interface {
	FetchHotel() ([]*models.Hotel, error)
}
