package user

import (
	"github.com/wincentrtz/bncc/models"
)

type Usecase interface {
	FetchUserById(userId int) (*models.User, error)
}
