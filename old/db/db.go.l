package db

import (
    "log"

    //"github.com/uvoo/minimon/pkg/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"fmt"
)

var DB *gorm.DB

// func Init(dbURL string) *gorm.DB {
func Init(dbURL string) {
    // dbURL := "postgres://pg:pass@localhost:5432/crud"

    DB, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }
	fmt.Printf("%v", DB)

    // db.AutoMigrate(&models.Book{})

    // return db
}
