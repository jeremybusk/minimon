package main

import (
	// "gorm.io/gorm"
	"fmt"
	"github.com/google/uuid"
	// "gorm.io/driver/postgres"
	"log"
	//"os"
	"time"
	// "uvoo.io/gormpg/v2/database"
	// database "database"
	// "github.com/satori/go.uuid"
	"github.com/uvoo/minimon/db"
	// "github.com/uvoo/minimon/pkg/hello"
	// "github.com/uvoo/minimon/pkg/db2"

	// "github.com/uvoo/minimon/pkg/models"
	"github.com/tcnksm/go-httpstat"
	"github.com/uvoo/minimon/models"
	// "github.com/shopspring/decimal"
	// "github.com/uvoo/minimon/pkg/handlers"

	"io"
	"io/ioutil"
	// "log"
	"net"
	"net/http"
	// "os"
	// "time"
)

func get_latency() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))

	start := time.Now()
	oneByte := make([]byte, 1)
	_, err = conn.Read(oneByte)
	if err != nil {
		panic(err)
	}
	log.Println("First byte:", time.Since(start))

	_, err = ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}
	log.Println("Everything:", time.Since(start))
}


func getURLStats2(url string) {
	// args := os.Args
	// if len(args) < 2 {
	//	log.Fatalf("Usage: go run main.go URL")
	//}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	result.End(time.Now())

	fmt.Printf("urlstats2: %+v\n", result)

}

func getURLStats(url string) {
	// Create a new HTTP request
	// req, err := http.NewRequest("GET", "https://github.com", nil)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Create a httpstat powered context
	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)
	// Send request by default HTTP client
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	end := time.Now()
	fmt.Printf("end %v", end)
	// Show the results
	log.Printf("DNS lookup: %d ms", int(result.DNSLookup/time.Millisecond))
	log.Printf("TCP connection: %d ms", int(result.TCPConnection/time.Millisecond))
	log.Printf("TLS handshake: %d ms", int(result.TLSHandshake/time.Millisecond))
	log.Printf("Server processing: %d ms", int(result.ServerProcessing/time.Millisecond))
	log.Printf("Content transfer: %d ms", int(result.ContentTransfer(time.Now())/time.Millisecond))
}

func main() {
	// dbURL := os.Getenv("MINIMON_DBURL")
	// var err error
	// db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
//     if err := db.Open(); err != nil {
       // handle error
// log.Fatal(err)
//     }
db.Open()
    // defer db.Close()
	// db.DBcon = db.Init(dbURL)
	dothis()


}

func dothis(){
	URL := models.URL{}
	path := "https://example.com"
	db.DB.First(&URL, "path = ?", path)
	db.DB.Model(&URL).Update("Rsp_time", 2)
	fmt.Printf("URL Path: %v\n", URL.Path)
	// fmt.Printf("URL Path: %v", tx)
	// tx := DB.First(&user, "name = ?", "Jeremy")
	getURLStats(path)

	path = "baseball.com"
	path = "gaga.com"
	getURLStats2(path)
}

// TRASH ===========================================
// tx := DB.First(&user, "name = ?", "Jeremy")
//	if tx.Error != nil {
//	    // fmt.Printf("%v", tx)
//		log.Fatal(tx)
//return false
//	}
// fmt.Printf("User: %v", &user)
// dataSourceName := "host=localhost user=minimon password=minimon dbname=gorm port=9999 sslmode=disable TimeZone=America/Denver"
// dbURL := "postgres://minimon:minimon@localhost:9999/gorm"
// dsn := "host=localhost user=minimon password=minimon dbname=gorm port=9999 sslmode=disable TimeZone=America/Denver"
// InitDB(dsn)
//d_init()
// db_create()
type FURL struct {
	URL_id       int64   `db:"url_id"`
	UUID         string  `db:"uuid"`
	URL_group_id int     `db:"url_group_id"`
	URL          string  `db:"url"`
	Note         string  `db:"note"`
	Rsp_regex    int     `db:"rsp_regex"`
	Rsp_code     int     `db:"rsp_code"`
	Sequence     int     `db:"sequence"`
	Disabled     bool    `db:"disabled"`
	Rsp_time     float64 `db:"rsp_time"`
	Created_at   string  `db:"created_at"`
	Updated_at   string  `db:"updated_at"`
	Completed_at string  `db:"completed_at"`
}

// Base contains common columns for all tables.
type fBase struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// type Profile struct {
// Base
// Name   string    `gorm:"column:name;size:128;not null;"`
// UserID uuid.UUID `gorm:"type:uuid;column:user_foreign_key;not null;"`
//}

type fNullString struct {
	String string
	Valid  bool // Valid is true if String is not NULL
}

// gorm.Model definition (This it not needed only for show)
// type Model struct {
//  ID        uint           `gorm:"primaryKey"`
//  CreatedAt time.Time
//  UpdatedAt time.Time
//j  DeletedAt gorm.DeletedAt `gorm:"index"`
// }

// ID           uint `gorm:"primaryKey"`
