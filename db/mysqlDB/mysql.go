package mysqlDB

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/viciousvs/go-tut-RestApi/models"
	"github.com/viciousvs/go-tut-RestApi/pkg/errors"
)

type Storage struct {
	DB *sql.DB
}

func (s *Storage) GetAllAlbum() ([]models.Album, error) {
	var albums = make([]models.Album, 0, 10)
	rows, err := s.DB.Query("SELECT * FROM album")
	if err != nil {
		// handle this error better than this
		return albums, &errors.NotFoundError{}
	}
	defer rows.Close()
	for rows.Next() {
		var alb models.Album
		err = rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
		if err != nil {
			// handle this error
			return albums, err
		}
		albums = append(albums, alb)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return albums, nil
}
func (s *Storage) CreateAlbum(alb models.Album) (int64, error) {
	result, err := s.DB.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
func (s *Storage) ReadAlbum(id int64) (models.Album, error) {
	var alb models.Album
	row := s.DB.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, &errors.NotFoundError{}
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}
func (s *Storage) DeleteAlbum(id int64) error {
	_, err := s.DB.Exec("DELETE FROM album WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
func (s *Storage) UpdateAlbum(album models.Album) error {
	return &errors.NotFoundError{}
}
