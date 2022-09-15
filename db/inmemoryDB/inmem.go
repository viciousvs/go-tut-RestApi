package inmemoryDB

import (
	"github.com/viciousvs/go-tut-RestApi/pkg/errors"

	"github.com/viciousvs/go-tut-RestApi/models"
)

type Storage struct {
	albums []models.Album
}

func (s *Storage) GetAllAlbum() ([]models.Album, error) {
	return s.albums, nil
}
func (s *Storage) CreateAlbum(album models.Album) (int64, error) {
	// Add the new album to the slice.
	s.albums = append(s.albums, album)
	return album.ID, nil
}
func (s *Storage) ReadAlbum(id int64) (models.Album, error) {
	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, album := range s.albums {
		if album.ID == id {
			return album, nil
		}
	}
	return models.Album{}, &errors.NotFoundError{}
}
func (s *Storage) DeleteAlbum(id int64) error {
	ind := -1
	for i, album := range s.albums {
		if album.ID == id {
			ind = i
			break
		}
	}

	if ind != -1 {
		s.albums = append(s.albums[:ind], s.albums[ind+1:]...)
		return nil
	}
	return &errors.NotFoundError{}
}
func (s *Storage) UpdateAlbum(album models.Album) error {
	for ind, a := range s.albums {
		if a.ID == album.ID {
			s.albums[ind] = album
			return nil
		}
	}
	return &errors.NotFoundError{}
}
