package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/wincentrtz/bncc/api/domain/hotel"
	"github.com/wincentrtz/bncc/api/models"
	"github.com/wincentrtz/bncc/api/models/builder"
)

type hotelRepository struct {
	Conn *sql.DB
}

func NewHotelRepository(Conn *sql.DB) hotel.Repository {
	return &hotelRepository{
		Conn,
	}
}

func (m *hotelRepository) FetchHotel() ([]*models.Hotel, error) {
	query := `
		SELECT * FROM hotel
		`
	rows, err := m.Conn.Query(query)
	defer rows.Close()
	if err != nil || rows == nil {
		fmt.Println(err)
		return nil, nil
	}

	hotels := make([]*models.Hotel, 0)
	for rows.Next() {
		var id int
		var name string
		var phone string
		var address string
		var city string
		var price int
		var created time.Time

		err = rows.Scan(
			&id,
			&name,
			&phone,
			&address,
			&city,
			&price,
			&created,
		)

		hotel := builder.NewHotel().
			Id(id).
			Name(name).
			Phone(phone).
			Address(address).
			City(city).
			Price(price).
			Created(created).
			Build()

		if err != nil {
			return nil, err
		}

		hotels = append(hotels, hotel)
	}

	return hotels, nil
}
