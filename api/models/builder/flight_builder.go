package builder

import (
	"time"

	"github.com/wincentrtz/bncc/api/models"
)

type flightBuilder struct {
	id          int       `json:"id"`
	date        time.Time `json:"date"`
	time        string    `json:"time"`
	departure   string    `json:"departure"`
	destination string    `json:"destination"`
	price       int       `json:"price"`
	created     time.Time `json:"created"`
}

type FlightBuilder interface {
	Id(id int) FlightBuilder
	Date(date time.Time) FlightBuilder
	Time(time string) FlightBuilder
	Departure(departure string) FlightBuilder
	Destination(destination string) FlightBuilder
	Price(price int) FlightBuilder
	Created(created time.Time) FlightBuilder
	Build() *models.Flight
}

func NewFlight() FlightBuilder {
	return &flightBuilder{}
}

func (fb *flightBuilder) Id(id int) FlightBuilder {
	fb.id = id
	return fb
}

func (fb *flightBuilder) Date(date time.Time) FlightBuilder {
	fb.date = date
	return fb
}

func (fb *flightBuilder) Time(time string) FlightBuilder {
	fb.time = time
	return fb
}

func (fb *flightBuilder) Departure(departure string) FlightBuilder {
	fb.departure = departure
	return fb
}

func (fb *flightBuilder) Destination(destination string) FlightBuilder {
	fb.destination = destination
	return fb
}

func (fb *flightBuilder) Price(price int) FlightBuilder {
	fb.price = price
	return fb
}

func (fb *flightBuilder) Created(created time.Time) FlightBuilder {
	fb.created = created
	return fb
}

func (fb *flightBuilder) Build() *models.Flight {
	return &models.Flight{
		Id:          fb.id,
		Date:        fb.date,
		Time:        fb.time,
		Departure:   fb.departure,
		Destination: fb.destination,
		Price:       fb.price,
		Created:     fb.created,
	}
}
