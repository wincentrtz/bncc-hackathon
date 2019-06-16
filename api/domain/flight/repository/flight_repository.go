package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/wincentrtz/bncc/api/domain/flight"
	"github.com/wincentrtz/bncc/api/models"
	"github.com/wincentrtz/bncc/api/models/builder"
)

type flightRepository struct {
	Conn *sql.DB
}

func NewFlightRepository(Conn *sql.DB) flight.Repository {
	return &flightRepository{
		Conn,
	}
}

func (m *flightRepository) FetchFlight() ([]*models.Flight, error) {
	query := `
		SELECT * FROM flight
		`
	rows, err := m.Conn.Query(query)
	defer rows.Close()
	if err != nil || rows == nil {
		fmt.Println(err)
		return nil, nil
	}

	flights := make([]*models.Flight, 0)
	for rows.Next() {
		var id int
		var date time.Time
		var departure string
		var destination string
		var price int
		var created time.Time
		var time string

		err = rows.Scan(
			&id,
			&date,
			&time,
			&departure,
			&destination,
			&price,
			&created,
		)

		flight := builder.NewFlight().
			Id(id).
			Date(date).
			Time(time).
			Departure(departure).
			Destination(destination).
			Price(price).
			Created(created).
			Build()

		if err != nil {
			return nil, err
		}

		flights = append(flights, flight)
	}

	return flights, nil
}
