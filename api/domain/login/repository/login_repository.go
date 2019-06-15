package repository

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"time"

	"github.com/wincentrtz/bncc/api/domain/login"
	"github.com/wincentrtz/bncc/api/models"
	"github.com/wincentrtz/bncc/api/models/builder"
	"github.com/wincentrtz/bncc/api/models/request"
)

type loginRepository struct {
	Conn *sql.DB
}

func NewLoginRepository(Conn *sql.DB) login.Repository {
	return &loginRepository{
		Conn,
	}
}

func (m *loginRepository) Login(lr *request.LoginRequest) (*models.Login, error) {
	var id int
	var name string
	var email string
	var password string
	var phone string
	var address string
	var gender string
	var ktp string
	var created time.Time

	hasher := md5.New()
	hasher.Write([]byte(lr.Password))
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

	query := `
		SELECT * FROM users
		WHERE 
			email = '` + lr.Username + `' AND
			password = '` + encryptedPassword + `'
		`

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

	login := builder.NewLogin().
		User(*user).
		Build()

	return login, nil
}
