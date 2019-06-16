package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/wincentrtz/bncc/api/domain/album"
	"github.com/wincentrtz/bncc/api/models"
	"github.com/wincentrtz/bncc/api/models/builder"
)

type albumRepository struct {
	Conn *sql.DB
}

func NewAlbumRepository(Conn *sql.DB) album.Repository {
	return &albumRepository{
		Conn,
	}
}

func (m *albumRepository) FetchAlbum() ([]*models.Album, error) {
	query := `
		SELECT * FROM albums
		`
	rows, err := m.Conn.Query(query)
	defer rows.Close()
	if err != nil || rows == nil {
		fmt.Println(err)
		return nil, nil
	}

	fmt.Println(rows)
	albums := make([]*models.Album, 0)
	for rows.Next() {
		var id int
		var name string
		var description string
		var isPaidOff int
		var created time.Time

		err = rows.Scan(
			&id,
			&name,
			&description,
			&isPaidOff,
			&created,
		)

		album := builder.NewAlbum().
			Id(id).
			Name(name).
			Description(description).
			IsPaidOff(isPaidOff).
			Created(created).
			Build()

		if err != nil {
			return nil, err
		}

		albums = append(albums, album)
	}

	return albums, nil
}
