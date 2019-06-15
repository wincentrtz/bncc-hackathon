package models

import "time"

// Payment models
type Payment struct {
	ID                     int              `json:"id"`
	Name                   string           `json:"name"`
	Description            string           `json:"description"`
	AlbumID                int              `json:"album_id"`
	TotalOutstandingAmount int              `json:"total_outstanding_amount"`
	TotalPaidAmount        int              `json:"total_paid_amount"`
	Details                []*PaymentDetail `json:"payment_details"`
	CreatedDate            time.Time        `json:"create_date"`
	DueDate                time.Time        `json:"due_date"`
	PaidOffDate            time.Time        `json:"paid_off_date"`
}
