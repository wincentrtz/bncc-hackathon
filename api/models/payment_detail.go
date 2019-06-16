package models

import "time"

// PaymentDetail models
type PaymentDetail struct {
	ID                int       `json:"id"`
	UserID            int       `json:"user_id"`
	PaymentID         int       `json:"payment_id"`
	OutstandingAmount int       `json:"outstanding_amount"`
	PaidAmount        int       `json:"paid_amount"`
	CreatedDate       time.Time `json:"create_date"`
	DueDate           time.Time `json:"due_date"`
	PaymentDate       time.Time `json:"payment_date"`
}
