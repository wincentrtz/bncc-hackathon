package builder

import (
	"time"

	"github.com/wincentrtz/bncc/api/models"
)

type hotelBuilder struct {
	id      int       `json:"id"`
	name    string    `json:"name"`
	phone   string    `json:"phone"`
	address string    `json:"address"`
	city    string    `json:"city"`
	price   int       `json:"price"`
	created time.Time `json:"created"`
}

type HotelBuilder interface {
	Id(id int) HotelBuilder
	Name(name string) HotelBuilder
	Phone(phone string) HotelBuilder
	Address(address string) HotelBuilder
	City(city string) HotelBuilder
	Price(price int) HotelBuilder
	Created(created time.Time) HotelBuilder
	Build() *models.Hotel
}

func NewHotel() HotelBuilder {
	return &hotelBuilder{}
}

func (hb *hotelBuilder) Id(id int) HotelBuilder {
	hb.id = id
	return hb
}

func (hb *hotelBuilder) Name(name string) HotelBuilder {
	hb.name = name
	return hb
}

func (hb *hotelBuilder) Phone(phone string) HotelBuilder {
	hb.phone = phone
	return hb
}

func (hb *hotelBuilder) Address(address string) HotelBuilder {
	hb.address = address
	return hb
}

func (hb *hotelBuilder) City(city string) HotelBuilder {
	hb.city = city
	return hb
}

func (hb *hotelBuilder) Price(price int) HotelBuilder {
	hb.price = price
	return hb
}

func (hb *hotelBuilder) Created(created time.Time) HotelBuilder {
	hb.created = created
	return hb
}

func (hb *hotelBuilder) Build() *models.Hotel {
	return &models.Hotel{
		Id:      hb.id,
		Name:    hb.name,
		Phone:   hb.phone,
		Address: hb.address,
		City:    hb.city,
		Price:   hb.price,
		Created: hb.created,
	}
}
