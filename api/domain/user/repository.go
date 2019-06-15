package user

import (
	"github.com/wincentrtz/bncc/api/models"
)

type Repository interface {
	FetchUserById(userId int) (*models.User, error)
}
