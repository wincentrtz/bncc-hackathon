package models

import "time"

type Flight struct {
	Id          int       `json:"id"`
	Date        time.Time `json:"date"`
	Time        string    `json:"time"`
	Departure   string    `json:"departure"`
	Destination string    `json:"destination"`
	Price       int       `json:"price"`
	Created     time.Time `json:"created"`
}
