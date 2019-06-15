package login

import (
	"github.com/wincentrtz/bncc/api/models"
	"github.com/wincentrtz/bncc/api/models/request"
)

type Repository interface {
	Login(lr *request.LoginRequest) (*models.Login, error)
}
