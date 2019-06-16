package user

import (
	"github.com/wincentrtz/bncc/api/models"
)

// Usecase interface
type Usecase interface {
	FetchUserById(userId int) (*models.User, error)
}
