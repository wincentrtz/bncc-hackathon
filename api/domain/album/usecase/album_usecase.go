package usecase

import (
	"time"

	"github.com/wincentrtz/bncc/api/domain/album"
	"github.com/wincentrtz/bncc/api/models"
)

type albumUsecase struct {
	albumRepo      album.Repository
	contextTimeout time.Duration
}

func NewAlbumUsecase(hr album.Repository, timeout time.Duration) album.Usecase {
	return &albumUsecase{
		albumRepo:      hr,
		contextTimeout: timeout,
	}
}

func (hu *albumUsecase) FetchAlbum() ([]*models.Album, error) {
	albums, err := hu.albumRepo.FetchAlbum()
	if err != nil {
		return nil, err
	}
	return albums, nil
}
