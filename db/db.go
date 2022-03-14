package db

import (
    "log"

    //"github.com/uvoo/minimon/pkg/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DBcon *gorm.DB

func Init(dbURL string) *gorm.DB {
    // dbURL := "postgres://pg:pass@localhost:5432/crud"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    // db.AutoMigrate(&models.Book{})

    return db
}
