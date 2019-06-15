package user

import (
	"github.com/wincentrtz/bncc/api/models"
)

type Usecase interface {
	FetchUserById(userId int) (*models.User, error)
}
