package builder

import (
	"github.com/wincentrtz/bncc/api/models"
)

type loginBuilder struct {
	user models.User
}

type LoginBuilder interface {
	User(models.User) LoginBuilder
	Build() *models.Login
}

func NewLogin() LoginBuilder {
	return &loginBuilder{}
}

func (lb *loginBuilder) User(user models.User) LoginBuilder {
	lb.user = user
	return lb
}

func (lb *loginBuilder) Build() *models.Login {
	return &models.Login{
		User: lb.user,
	}
}
