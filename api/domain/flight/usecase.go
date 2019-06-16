package flight

import (
	"github.com/wincentrtz/bncc/api/models"
)

type Usecase interface {
	FetchFlight() ([]*models.Flight, error)
}
