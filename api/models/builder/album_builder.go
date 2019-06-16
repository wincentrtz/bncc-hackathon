package builder

import (
	"time"

	"github.com/wincentrtz/bncc/api/models"
)

type albumBuilder struct {
	id          int
	name        string
	description string
	isPaidOff   int
	created     time.Time
}

type AlbumBuilder interface {
	Id(int) AlbumBuilder
	Name(string) AlbumBuilder
	Description(string) AlbumBuilder
	IsPaidOff(int) AlbumBuilder
	Created(time.Time) AlbumBuilder
	Build() *models.Album
}

// NewAlbum factory function
func NewAlbum() AlbumBuilder {
	return &albumBuilder{}
}

func (ab *albumBuilder) Id(id int) AlbumBuilder {
	ab.id = id
	return ab
}

func (ab *albumBuilder) Name(name string) AlbumBuilder {
	ab.name = name
	return ab
}

func (ab *albumBuilder) Description(description string) AlbumBuilder {
	ab.description = description
	return ab
}

func (ab *albumBuilder) IsPaidOff(isPaidOff int) AlbumBuilder {
	ab.isPaidOff = isPaidOff
	return ab
}

func (ab *albumBuilder) Created(created time.Time) AlbumBuilder {
	ab.created = created
	return ab
}

func (ab *albumBuilder) Build() *models.Album {
	return &models.Album{
		ID:          ab.id,
		Name:        ab.name,
		Description: ab.description,
		IsPaidOff:   ab.isPaidOff,
		Created:     ab.created,
	}
}
