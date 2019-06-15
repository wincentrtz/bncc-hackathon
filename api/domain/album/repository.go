package album

import (
	"database/sql"

	"github.com/wincentrtz/bncc/api/models"
)

// Repository interface
type Repository interface {
	FetchAll() ([]*models.Album, error)
	FetchAlbumByID(albumID int) (*models.Album, error)
	CreateAlbum(name string, description string, users []*models.User, imagePaths []string) (*models.Album, error)
}

type albumRepository struct {
	Conn *sql.DB
}

// NewAlbumRepository factory
func NewAlbumRepository(Conn *sql.DB) Repository {
	return &albumRepository{
		Conn,
	}
}

func (ar *albumRepository) FetchAll() ([]*models.Album, error) {
	var albumPath = make(map[int][]string)
	var albumMaps = make(map[int]*models.Album)
	var albums = []*models.Album{}

	query := `
		SELECT a.id, a.name, a.description, a.isPaidOff, ip.imagePath
		FROM albums AS a
		JOIN imagePaths AS ip ON a.id = ip.album_id`

	rows, err := ar.Conn.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		album := &models.Album{}
		var imagePath string

		rows.Scan(&album.ID, &album.Name, &album.Description, &album.IsPaidOff, &imagePath)

		albumMaps[album.ID] = album

		if ap, ok := albumPath[album.ID]; ok {
			albumPath[album.ID] = append(ap, imagePath)
		} else {
			albumPath[album.ID] = []string{imagePath}
		}
	}

	for albumID, imagePaths := range albumPath {
		var album *models.Album
		album = albumMaps[albumID]
		album.ImagePaths = imagePaths
		albums = append(albums, album)
	}

	return albums, nil
}

func (ar *albumRepository) FetchAlbumByID(albumID int) (*models.Album, error) {
	return nil, nil
}

func (ar *albumRepository) CreateAlbum(name string, description string, users []*models.User, imagePaths []string) (*models.Album, error) {
	return nil, nil
}
