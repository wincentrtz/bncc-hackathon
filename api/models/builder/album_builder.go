package builder

import "github.com/wincentrtz/bncc/api/models"

type albumBuilder struct {
	ID          int
	Name        string
	Description string
	Users       []*models.User
	Payments    []*models.Payment
	IsPaidOff   bool
	ImagePaths  []string
}

// AlbumBuilder interface
type AlbumBuilder interface {
	SetID(int) AlbumBuilder
	SetName(string) AlbumBuilder
	SetDescription(string) AlbumBuilder
	SetUsers([]*models.User) AlbumBuilder
	SetPayments([]*models.Payment) AlbumBuilder
	SetIsPaidOff(bool) AlbumBuilder
	SetImagePaths([]string) AlbumBuilder
	Build() *models.Album
}

// NewAlbum factory function
func NewAlbum() AlbumBuilder {
	return &albumBuilder{}
}

func (ab *albumBuilder) SetID(id int) AlbumBuilder {
	ab.ID = id
	return ab
}

func (ab *albumBuilder) SetName(name string) AlbumBuilder {
	ab.Name = name
	return ab
}

func (ab *albumBuilder) SetDescription(desc string) AlbumBuilder {
	ab.Description = desc
	return ab
}

func (ab *albumBuilder) SetUsers(users []*models.User) AlbumBuilder {
	ab.Users = users
	return ab
}

func (ab *albumBuilder) SetPayments(payments []*models.Payment) AlbumBuilder {
	ab.Payments = payments
	return ab
}

func (ab *albumBuilder) SetIsPaidOff(paidOff bool) AlbumBuilder {
	ab.IsPaidOff = paidOff
	return ab
}

func (ab *albumBuilder) SetImagePaths(paths []string) AlbumBuilder {
	ab.ImagePaths = paths
	return ab
}

func (ab *albumBuilder) Build() *models.Album {
	return &models.Album{
		ID:          ab.ID,
		Name:        ab.Name,
		Description: ab.Description,
		Users:       ab.Users,
		Payments:    ab.Payments,
		IsPaidOff:   ab.IsPaidOff,
		ImagePaths:  ab.ImagePaths,
	}
}
