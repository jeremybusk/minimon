package db

import (
    "log"

    //"github.com/uvoo/minimon/pkg/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	// "fmt"
"os"
)

var DB *gorm.DB

// func Open() *gorm.DB {
func Open() {
   var err error
   // DB, err = gorm.Open("postgres", "host=localhost port=5432 user=someUser dbname=someDB password=somePW sslmode=disable")
   dbURL := os.Getenv("MINIMON_DBURL")
    DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
   if err != nil {
       log.Fatalln(err)
       // return err
   }
   // fmt.Printf("%v", DB)
   // return DB
}
