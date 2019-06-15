package user

import (
	"github.com/wincentrtz/bncc/models"
)

type Repository interface {
	FetchUserById(userId int) (*models.User, error)
}
