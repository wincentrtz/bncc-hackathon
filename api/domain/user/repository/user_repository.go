package repository

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/wincentrtz/bncc/api/domain/user"
	"github.com/wincentrtz/bncc/api/models"
	"github.com/wincentrtz/bncc/api/models/builder"
)

type userRepository struct {
	Conn *sql.DB
}

func NewUserRepository(Conn *sql.DB) user.Repository {
	return &userRepository{
		Conn,
	}
}

func (m *userRepository) FetchUserById(userId int) (*models.User, error) {
	var id int
	var name string
	var email string
	var password string
	var phone string
	var address string
	var gender string
	var ktp string
	var created time.Time

	query := `
		SELECT * FROM users
		WHERE id =` + strconv.Itoa(userId)

	err := m.Conn.QueryRow(query).Scan(
		&id,
		&name,
		&email,
		&password,
		&phone,
		&address,
		&gender,
		&ktp,
		&created,
	)

	if err != nil {
		return nil, err
	}

	user := builder.NewUser().
		Id(id).
		Name(name).
		Email(email).
		Phone(phone).
		Address(address).
		Gender(gender).
		Ktp(ktp).
		Created(created).
		Build()

	return user, nil
}
