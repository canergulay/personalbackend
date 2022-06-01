package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PGManager struct {
	db *gorm.DB
}

func InitPG() PGManager {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		HOST,
		USERNAME,
		PASSWORD,
		DB,
		PORT,
		SSL)

	fmt.Println("WORKED I GUESS", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("AN UNEXPECTED ERROR !", err)
	}
	return PGManager{db: db}
}
