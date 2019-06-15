package login

import (
	"github.com/wincentrtz/bncc/api/models"
	"github.com/wincentrtz/bncc/api/models/request"
)

type Usecase interface {
	Login(lr *request.LoginRequest) (*models.Login, error)
}
