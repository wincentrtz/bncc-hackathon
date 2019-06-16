package album

import (
	"github.com/wincentrtz/bncc/api/models"
)

type Usecase interface {
	FetchAlbum() ([]*models.Album, error)
}
