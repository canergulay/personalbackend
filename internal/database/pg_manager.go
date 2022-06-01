package database

import (
	"fmt"
	"log"

	"github.canergulay/blogbackend/internal/server/endpoints/routes/blog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PGManager struct {
	DB *gorm.DB
}

func InitPG() PGManager {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		HOST,
		USERNAME,
		PASSWORD,
		DB,
		PORT,
		SSL)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	autoMigrater(db)
	if err != nil {
		log.Fatalln("AN UNEXPECTED ERROR !", err)
	}
	return PGManager{DB: db}
}

func autoMigrater(db *gorm.DB) {
	db.AutoMigrate(&blog.Post{})
}
