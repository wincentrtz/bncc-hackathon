package models

import "time"

type Album struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"desc"`
	IsPaidOff   int       `json:"is_paid_off"`
	Created     time.Time `json:"created"`
}
