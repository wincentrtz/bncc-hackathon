package album

import (
	"time"

	"github.com/wincentrtz/bncc/api/models"
)

// Usecase interface
type Usecase interface {
	FetchAll() ([]*models.Album, error)
	FetchAlbumByID(albumID int) (*models.Album, error)
	CreateAlbum(name string, description string, users []*models.User, imagePaths []string) (*models.Album, error)
}

type albumUsecase struct {
	albumRepo      Repository
	contextTimeout time.Duration
}

// NewAlbumUsecase factory
func NewAlbumUsecase(al Repository, timeout time.Duration) Usecase {
	return &albumUsecase{
		albumRepo:      al,
		contextTimeout: timeout,
	}
}

func (au *albumUsecase) FetchAll() ([]*models.Album, error) {
	albums, err := au.albumRepo.FetchAll()
	if err != nil {
		return nil, err
	}
	return albums, nil
}

func (au *albumUsecase) FetchAlbumByID(albumID int) (*models.Album, error) {
	return nil, nil
}

func (au *albumUsecase) CreateAlbum(name string, description string, users []*models.User, imagePaths []string) (*models.Album, error) {
	return nil, nil
}
