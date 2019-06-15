package builder

import (
	"time"

	"github.com/wincentrtz/bncc/api/models"
)

type paymentDetailBuilder struct {
	ID                int
	UserID            int
	PaymentID         int
	OutstandingAmount int
	PaidAmount        int
	CreatedDate       time.Time
	DueDate           time.Time
	PaymentDate       time.Time
}

// PaymentDetailBuilder interface
type PaymentDetailBuilder interface {
	SetID(id int) PaymentDetailBuilder
	SetUserID(userID int) PaymentDetailBuilder
	SetPaymentID(paymentID int) PaymentDetailBuilder
	SetOutstandingAmount(outsAmount int) PaymentDetailBuilder
	SetPaidAmount(paidAmount int) PaymentDetailBuilder
	SetCreatedDate(createdDate time.Time) PaymentDetailBuilder
	SetDueDate(dueDate time.Time) PaymentDetailBuilder
	SetPaymentDate(paidOffDate time.Time) PaymentDetailBuilder
	Build() *models.PaymentDetail
}

// NewPaymentDetail factory function
func NewPaymentDetail() PaymentDetailBuilder {
	return &paymentDetailBuilder{}
}

func (pdb *paymentDetailBuilder) SetID(id int) PaymentDetailBuilder {
	pdb.ID = id
	return pdb
}

func (pdb *paymentDetailBuilder) SetUserID(userID int) PaymentDetailBuilder {
	pdb.UserID = userID
	return pdb
}

func (pdb *paymentDetailBuilder) SetPaymentID(paymentID int) PaymentDetailBuilder {
	pdb.PaymentID = paymentID
	return pdb
}

func (pdb *paymentDetailBuilder) SetOutstandingAmount(outsAmount int) PaymentDetailBuilder {
	pdb.OutstandingAmount = outsAmount
	return pdb
}

func (pdb *paymentDetailBuilder) SetPaidAmount(paidAmount int) PaymentDetailBuilder {
	pdb.PaidAmount = paidAmount
	return pdb
}

func (pdb *paymentDetailBuilder) SetCreatedDate(createdDate time.Time) PaymentDetailBuilder {
	pdb.CreatedDate = createdDate
	return pdb
}

func (pdb *paymentDetailBuilder) SetDueDate(dueDate time.Time) PaymentDetailBuilder {
	pdb.DueDate = dueDate
	return pdb
}

func (pdb *paymentDetailBuilder) SetPaymentDate(paymentDate time.Time) PaymentDetailBuilder {
	pdb.PaymentDate = paymentDate
	return pdb
}

func (pdb *paymentDetailBuilder) Build() *models.PaymentDetail {
	return &models.PaymentDetail{
		ID:                pdb.ID,
		UserID:            pdb.UserID,
		PaymentID:         pdb.PaymentID,
		OutstandingAmount: pdb.OutstandingAmount,
		PaidAmount:        pdb.PaidAmount,
		CreatedDate:       pdb.CreatedDate,
		DueDate:           pdb.DueDate,
		PaymentDate:       pdb.PaymentDate,
	}
}
