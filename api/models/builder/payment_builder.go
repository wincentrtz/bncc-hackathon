package builder

import (
	"time"

	"github.com/wincentrtz/bncc/api/models"
)

type paymentBuilder struct {
	ID                     int
	Name                   string
	Description            string
	AlbumID                int
	TotalOutstandingAmount int
	TotalPaidAmount        int
	CreatedDate            time.Time
	DueDate                time.Time
	PaidOffDate            time.Time
}

// PaymentBuilder interface
type PaymentBuilder interface {
	SetID(id int) PaymentBuilder
	SetName(name string) PaymentBuilder
	SetDescription(desc string) PaymentBuilder
	SetAlbumID(albumID int) PaymentBuilder
	SetTotalOutstandingAmount(outsAmount int) PaymentBuilder
	SetTotalPaidAmount(paidAmount int) PaymentBuilder
	SetCreatedDate(createdDate time.Time) PaymentBuilder
	SetDueDate(dueDate time.Time) PaymentBuilder
	SetPaidOffDate(paidOffDate time.Time) PaymentBuilder
	Build() *models.Payment
}

// NewPayment factory function
func NewPayment() PaymentBuilder {
	return &paymentBuilder{}
}

func (pb *paymentBuilder) SetID(id int) PaymentBuilder {
	pb.ID = id
	return pb
}

func (pb *paymentBuilder) SetName(name string) PaymentBuilder {
	pb.Name = name
	return pb
}

func (pb *paymentBuilder) SetDescription(desc string) PaymentBuilder {
	pb.Description = desc
	return pb
}

func (pb *paymentBuilder) SetAlbumID(albumID int) PaymentBuilder {
	pb.AlbumID = albumID
	return pb
}

func (pb *paymentBuilder) SetTotalOutstandingAmount(outsAmount int) PaymentBuilder {
	pb.TotalOutstandingAmount = outsAmount
	return pb
}

func (pb *paymentBuilder) SetTotalPaidAmount(paidAmount int) PaymentBuilder {
	pb.TotalPaidAmount = paidAmount
	return pb
}

func (pb *paymentBuilder) SetCreatedDate(createdDate time.Time) PaymentBuilder {
	pb.CreatedDate = createdDate
	return pb
}

func (pb *paymentBuilder) SetDueDate(dueDate time.Time) PaymentBuilder {
	pb.DueDate = dueDate
	return pb
}

func (pb *paymentBuilder) SetPaidOffDate(paidOffDate time.Time) PaymentBuilder {
	pb.PaidOffDate = paidOffDate
	return pb
}

func (pb *paymentBuilder) Build() *models.Payment {
	return &models.Payment{
		ID:                     pb.ID,
		Name:                   pb.Name,
		Description:            pb.Description,
		AlbumID:                pb.AlbumID,
		TotalOutstandingAmount: pb.TotalOutstandingAmount,
		TotalPaidAmount:        pb.TotalPaidAmount,
		CreatedDate:            pb.CreatedDate,
		DueDate:                pb.DueDate,
		PaidOffDate:            pb.PaidOffDate,
	}
}
