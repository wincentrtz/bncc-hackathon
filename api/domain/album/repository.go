package album

import (
	"github.com/wincentrtz/bncc/api/models"
)

type Repository interface {
	FetchAlbum() ([]*models.Album, error)
}
