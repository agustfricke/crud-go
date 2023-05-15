package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {

  err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dbDNS := os.Getenv("DB_DNS")

  var error error

  DB, error = gorm.Open(postgres.Open(dbDNS), &gorm.Config{})

  if error != nil {
    log.Fatal(error)
  } else {
    log.Println("Database connected")
  }
}

