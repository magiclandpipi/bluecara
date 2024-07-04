package database

import (
    "fmt"
    "log"
    "os"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Connect() {
    var err error
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

    DB, err = gorm.Open("postgres", dsn)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    DB.AutoMigrate(&models.Booking{})
}
