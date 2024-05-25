package db

import (
	"go-ddd-practice/pkg/model"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func DbInit() *gorm.DB {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	db, err = gorm.Open(mysql.Open(os.Getenv("DB_PATH")))
	if err != nil {
		log.Fatal(err)
	}

	autoMigrate()

	return db
}

func GetDB() *gorm.DB {
	return db
}

func autoMigrate() {
	db.AutoMigrate(
		model.Accounts{},
	)
}
