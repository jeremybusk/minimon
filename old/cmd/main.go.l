package main

import (
	"gorm.io/gorm"
	// "gorm.io/driver/sqlite"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"time"
	"log"
	// "uvoo.io/gormpg/v2/database"
	// database "database"
	// "github.com/satori/go.uuid"
	"github.com/uvoo/minimon/pkg/db"
	"github.com/tcnksm/go-httpstat"
	// "github.com/shopspring/decimal"
	// "github.com/uvoo/minimon/pkg/handlers"

	"io"
	"io/ioutil"
	// "log"
	"net/http"
    "net"
	// "os"
	// "time"
)


//type handler struct {
//    DB *gorm.DB
//}

//func New(db *gorm.DB) handler {
//    return handler{db}
//}

type User struct {
	gorm.Model
	UUID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	// ID           uint
	Name         string
	Age          uint8
	Birthday     time.Time
	Foo          string
	Foo2         string
	Email        string
	ActivatedAt  sql.NullTime
	MemberNumber sql.NullString
}

// Birthday     *time.Time
// Something:   sql.NullString{String: "a string goes here", Valid: true},

type URL struct {
	gorm.Model
	//URL_id       int64  `gorm:"primaryKey"`
	Disabled     bool
	UUID         string `gorm:"type:uuid;default:uuid_generate_v4()"`
	Note         string
	URL_group_id int
	Path          string `gorm:"unique;not null"`
	Rsp_code     int
	Rsp_code_exp int `gorm:"default:200`
	Rsp_code_test   bool
	Rsp_time     float64 `gorm:"type:decimal(16,6);default:0"`
	Rsp_time_exp int `gorm:"default:4`
	Rsp_time_test   bool
	Rsp_regex_exp string `gorm:"default:statushealthy`
	Rsp_regex_test   bool
	AllowInsecureTLS	bool `gorm:"default:false`
	// Rsp_time     Decimal `gorm:"type:decimal(16,6);default:0"`
	//Amount       float32   `sql:"type:decimal(10,2);"` 
	Sequence     int
	// Rsp_time     float64
}

// var dsn string = "host=localhost user=minimon password=minimon dbname=gorm port=9999 sslmode=disable TimeZone=America/Denver"
// var db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})


// var dsn string
// var db *gorm.DB
var DB *gorm.DB

// func InitDB(dataSourceName string) error {
// func InitDB(dataSourceName string) error {
func InitDB(dsn string)  *gorm.DB {
    // var err error

    // db, err = sql.Open("postgres", dataSourceName)
	// dsn := "host=localhost user=minimon password=minimon dbname=gorm port=9999 sslmode=disable TimeZone=America/Denver"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
		log.Fatalln(err)
    }
	// db.AutoMigrate(&models.Book{})
	// db.AutoMigrate(&users{})

    // return db.Ping()
    return DB
}
func db_create() {

	dsn := "host=localhost user=minimon password=minimon dbname=gorm port=9999 sslmode=disable TimeZone=America/Denver"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&User{})
	// db.AutoMigrate(&Profile{})

	// db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	user := User{Name: "Jeremy", Age: 45, Birthday: time.Now()}

	result := db.Create(&user) // pass pointer of data to Create

	// user.ID             // returns inserted data's primary key
	// user.Name             // returns inserted data's primary key
	fmt.Println(user.Name)
	return
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
	fmt.Println(err)
	// result.Error        // returns error
	// result.RowsAffected // returns inserted records count
	db.Take(&result)
}

func GetUser(DB *gorm.DB) {
// var DB *gorm.DB
	// dsn := "host=localhost user=minimon password=minimon dbname=gorm port=9999 sslmode=disable TimeZone=America/Denver"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	//     fmt.Printf("ERROR")
// 		log.Fatal(err)
// 	}
	// var user User
	user := User{}
	// var users []User
// tx := db.Take(&user)
// tx := db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
// tx := db.First(&user, "name = 'Jeremy'")
 // tx := db.First(&user, "name = ?", "Jeremy")
 tx := DB.First(&user, "name = ?", "Jeremy")
	if tx.Error != nil {
	    // fmt.Printf("%v", tx)
		log.Fatal(tx)
	    //return false
	}
	    fmt.Printf("Name: %+v\n", user.Name)
	    fmt.Printf("Name: %v\n", user.Name)
	    fmt.Printf("User: %#v\n", user)
	    fmt.Printf("%v", tx)
	    fmt.Printf("Name: %v", tx.Name)
	    fmt.Printf("DONE")
}

func get_latency(){
conn, err := net.Dial("tcp", "google.com:80")
if err != nil {
    panic(err)
}
defer conn.Close()
conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))

start := time.Now()
oneByte := make([]byte,1)
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

func d_init() {
	// db, err := gorm.Open(postgres.Open("gorm.db"), &gorm.Config{})
	dsn := "host=localhost user=minimon password=minimon dbname=gorm port=9999 sslmode=disable TimeZone=America/Denver"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	user := "jeremy1"

	// Most CRUD API kept compatibility
	// db.AutoMigrate(&Product{})
	fmt.Println(err)
	// db.First(&user, 1)
	// db.Raw("SELECT name FROM foo WHERE name = 'jeremy1'", 3).Scan(&result)
	var name string
	// db.Raw("SELECT name FROM users WHERE name = 'jeremy1'").Scan(&name)
	// var exists bool
	var exists bool
	sqlStr := "SELECT name FROM users WHERE name = 'Jeremy'"
	// db.Raw("SELECT name FROM users WHERE name = 'Jeremy'").Scan(&name)
	tx := db.Raw(sqlStr).Row().Scan(&exists)
	// tx := DB.Exec(sqlStr, args...)
	if tx.Error != nil {
	    fmt.Printf("ERROR")
	    // fmt.Printf("%v", tx.Error)
	    fmt.Printf("%v", tx)
		log.Fatal(err)
	    //return false
	}
	fmt.Printf("result: %v", tx)
	fmt.Printf("exists: %v", exists)
	println("done")
	fmt.Printf("name: %v", name)
	// println(&name)
	// println(*name)
	println("done")
	return
	// db.Create(&user)
	return
	db.First(&user, 1)
	db.Model(&user).Update("Age", 18)
	db.Model(&user).Omit("Role").Updates(map[string]interface{}{"Name": "jinzhu", "Role": "admin"})
	db.Delete(&user)
}


func getURLStats2(url string){
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

func getURLStats(url string){
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
dbURL := "postgres://minimon:minimon@localhost:9999/gorm"
DB := db.Init(dbURL)
DB.AutoMigrate(
	&User{},
	&URL{} )
	// user := User{}
// fmt.Printf("%v", DB)
	 // GetUser(DB)
	 // db.Create(&URL{Path: "D42", Price: 100})
	 // DB.Create(&URL{Path: "https://example.com"})
	 URL := URL{}
	 path := "https://example.com"
	 DB.First(&URL, "path = ?", path)
	 DB.Model(&URL).Update("Rsp_time", 2)
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
