package models

import "time"

type User struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Phone   string    `json:"phone"`
	Address string    `json:"address"`
	Gender  string    `json:"gender"`
	Ktp     string    `json:"ktp"`
	Created time.Time `json:"created_on"`
}
