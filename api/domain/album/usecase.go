package album

import (
	"github.com/wincentrtz/bncc/api/models"
)

// Usecase interface
type Usecase interface {
	FetchAll() ([]*models.Album, error)
	FetchAlbumByID(albumID int) (*models.Album, error)
	CreateAlbum(name string, description string, users []*models.User, imagePaths []string)
}
