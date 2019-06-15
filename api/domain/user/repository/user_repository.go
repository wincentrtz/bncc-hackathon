package repository

import (
	"database/sql"
	"strconv"

	"github.com/wincentrtz/bncc/domain/user"
	"github.com/wincentrtz/bncc/models"
	"github.com/wincentrtz/bncc/models/builder"
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

	query := `
		SELECT id, name FROM users
		WHERE id =` + strconv.Itoa(userId)

	err := m.Conn.QueryRow(query).Scan(
		&id,
		&name,
	)

	if err != nil {
		return nil, err
	}

	user := builder.NewUser().
		Id(id).
		Name(name).
		Build()

	return user, nil
}
