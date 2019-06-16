package flight

import (
	"github.com/wincentrtz/bncc/api/models"
)

type Repository interface {
	FetchFlight() ([]*models.Flight, error)
}
