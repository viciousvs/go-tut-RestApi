package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/viciousvs/go-tut-RestApi/models"
	"github.com/viciousvs/go-tut-RestApi/pkg/errors"
)

// getAlbums responds with the list of all albums as JSON.
func (s *Service) getAlbums(c *gin.Context) {
	albums, err := s.Storage.GetAllAlbum()
	if err != nil {
		code, msg := errors.CheckError(err)
		c.IndentedJSON(code, msg)
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func (s *Service) postAlbums(c *gin.Context) {
	var newAlbum models.Album
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	id, err := s.Storage.CreateAlbum(newAlbum)
	if err != nil {
		code, msg := errors.CheckError(err)
		c.IndentedJSON(code, msg)
		return
	}
	newAlbum.ID = id
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (s *Service) getAlbumByID(c *gin.Context) {
	pid := c.Param("id")
	id, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		return
	}
	album, err := s.Storage.ReadAlbum(id)
	if err != nil {
		code, msg := errors.CheckError(err)
		c.IndentedJSON(code, msg)
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func (s *Service) deleteById(c *gin.Context) {
	pid := c.Param("id")
	id, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		return
	}
	if err := s.Storage.DeleteAlbum(id); err != nil {
		code, msg := errors.CheckError(err)
		c.IndentedJSON(code, msg)
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "album deleted"})
}
