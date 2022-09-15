package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/viciousvs/go-tut-RestApi/db/mysqlDB"
	"github.com/viciousvs/go-tut-RestApi/models"
)

type AlbumStorage interface {
	GetAllAlbum() ([]models.Album, error)
	CreateAlbum(album models.Album) (int64, error)
	ReadAlbum(id int64) (models.Album, error)
	DeleteAlbum(id int64) error
	UpdateAlbum(album models.Album) error
}
type Storage struct {
	AlbumStorage
}
type Service struct {
	Storage Storage
}

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	router := gin.Default()
	// inmemstorage := Storage{
	// 	&inmemoryDB.Storage{},
	// }
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               os.Getenv("DBNAME"),
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("DB Connected")

	storage := Storage{
		&mysqlDB.Storage{DB: db},
	}
	s := Service{
		Storage: storage,
	}
	router.GET("/albums", s.getAlbums)
	router.GET("/albums/:id", s.getAlbumByID)
	router.POST("/albums", s.postAlbums)
	router.DELETE("/albums/:id", s.deleteById)

	router.Run("localhost:8080")
}
