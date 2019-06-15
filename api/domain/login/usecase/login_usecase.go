package usecase

import (
	"time"

	"github.com/wincentrtz/bncc/api/domain/login"
	"github.com/wincentrtz/bncc/api/models"
	"github.com/wincentrtz/bncc/api/models/request"
)

type loginUsecase struct {
	loginRepo      login.Repository
	contextTimeout time.Duration
}

func NewLoginUsecase(lr login.Repository, timeout time.Duration) login.Usecase {
	return &loginUsecase{
		loginRepo:      lr,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) Login(lr *request.LoginRequest) (*models.Login, error) {
	user, err := lu.loginRepo.Login(lr)
	if err != nil {
		return nil, err
	}
	return user, nil
}
