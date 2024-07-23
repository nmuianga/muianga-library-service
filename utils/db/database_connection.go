package db

import (
	"fmt"
	"github.com/nmuianga/muianga-library-service/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func OpenConnection() {
	dsn := "host=localhost user=postgres password=postgres dbname=muianga_library port=5432 sslmode=disable TimeZone=Africa/Maputo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	err = db.AutoMigrate(&model.BookCategory{}, &model.Publisher{}, &model.Book{})
	if err != nil {
		log.Fatal("There was an error connecting to database")
	}
}
