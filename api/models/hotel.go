package models

import "time"

type Hotel struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Phone   string    `json:"phone"`
	Address string    `json:"address"`
	City    string    `json:"city"`
	Price   int       `json:"price"`
	Created time.Time `json:"create"`
}
