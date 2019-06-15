package builder

import (
	"time"

	"github.com/wincentrtz/bncc/api/models"
)

type userBuilder struct {
	id      int
	name    string
	email   string
	phone   string
	address string
	gender  string
	ktp     string
	created time.Time
}

type UserBuilder interface {
	Id(int) UserBuilder
	Name(string) UserBuilder
	Email(string) UserBuilder
	Phone(string) UserBuilder
	Address(string) UserBuilder
	Gender(string) UserBuilder
	Ktp(string) UserBuilder
	Created(time.Time) UserBuilder
	Build() *models.User
}

func NewUser() UserBuilder {
	return &userBuilder{}
}

func (ub *userBuilder) Id(id int) UserBuilder {
	ub.id = id
	return ub
}

func (ub *userBuilder) Email(email string) UserBuilder {
	ub.email = email
	return ub
}

func (ub *userBuilder) Name(name string) UserBuilder {
	ub.name = name
	return ub
}

func (ub *userBuilder) Phone(phone string) UserBuilder {
	ub.phone = phone
	return ub
}

func (ub *userBuilder) Address(address string) UserBuilder {
	ub.address = address
	return ub
}

func (ub *userBuilder) Gender(gender string) UserBuilder {
	ub.gender = gender
	return ub
}

func (ub *userBuilder) Ktp(ktp string) UserBuilder {
	ub.ktp = ktp
	return ub
}

func (ub *userBuilder) Created(created time.Time) UserBuilder {
	ub.created = created
	return ub
}

func (ub *userBuilder) Build() *models.User {
	return &models.User{
		Id:      ub.id,
		Email:   ub.email,
		Name:    ub.name,
		Phone:   ub.phone,
		Address: ub.address,
		Gender:  ub.gender,
		Ktp:     ub.ktp,
		Created: ub.created,
	}
}
