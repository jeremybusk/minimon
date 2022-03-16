package models

import (
	// "database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/tcnksm/go-httpstat"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"log"
	"minimon/database"
	"net/http"
	"os"
	"time"
	"database/sql"
	"github.com/jackc/pgtype"
	// https://github.com/davecheney/httpstat.git
	// "time"
)

// sql.NullString
// Int64 32 16 Time
type Platform struct {
	gorm.Model
	Disabled         bool `gorm:"type:bool;default:false"`
	UUID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Note             sql.NullString
	Sequence int     `gorm:"type:int;default:0"`
	Name             sql.NullString 
}


type Group struct {
	gorm.Model
	//URL_id       int64  `gorm:"primaryKey"`
	Disabled        bool `gorm:"type:bool;default:false"` 
	UUID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Note            sql.NullString
	name string
}


//fmt.Println(namelookup, connect, pretransfer, starttransfer, total)
type HTTPConnectionTrigger struct {
	gorm.Model
	UUID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
    DomainNameLookupTime float64 `gorm:"type:decimal(16,6);default:500"` //:= t1.Sub(t0)
    TCPConnectionTime float64 `gorm:"type:decimal(16,6);default:500"` //:= t2.Sub(t1)
    ConnectTime float64 `gorm:"type:decimal(16,6);default:500"` //:= t2.Sub(t0)
    PreTransferTime float64 `gorm:"type:decimal(16,6);default:500"` //:= t3.Sub(t0)
    StartTransferTime float64 `gorm:"type:decimal(16,6);default:500"` //:= t4.Sub(t0)
    ServerProcessingTime float64 `gorm:"type:decimal(16,6);default:500"` //:= t4.Sub(t3)
    TLSHandshakeTime float64 `gorm:"type:decimal(16,6);default:500"` //:= t6.Sub(t5)
    ContextTransferTime  float64 `gorm:"type:decimal(16,6);default:500"`//:= t7.Sub(t4)
    TotalTime float64 `gorm:"type:decimal(16,6);default:10000"` //:= t7.Sub(t0)
	StatusCode int16 `gorm:"default:200`
	TextRegexMatch string  `gorm:"default:megamonstatushealthy`
    IPAddress  pgtype.Inet    `gorm:"type:inet;default:'0.0.0.0/0'"`
}

//fmt.Println(namelookup, connect, pretransfer, starttransfer, total)
type HTTPConnection struct {
	gorm.Model
	UUID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	URLID int
    DomainNameLookupTime float64 `gorm:"type:decimal(16,6);default:0"` //:= t1.Sub(t0)
    TCPConnectionTime float64 `gorm:"type:decimal(16,6);default:0"` //:= t2.Sub(t1)
    ConnectTime float64 `gorm:"type:decimal(16,6);default:0"` //:= t2.Sub(t0)
    PreTransferTime float64 `gorm:"type:decimal(16,6);default:0"` //:= t3.Sub(t0)
    StartTransferTime float64 `gorm:"type:decimal(16,6);default:0"` //:= t4.Sub(t0)
    ServerProcessingTime float64 `gorm:"type:decimal(16,6);default:0"` //:= t4.Sub(t3)
    TLSHandshakeTime float64 `gorm:"type:decimal(16,6);default:0"` //:= t6.Sub(t5)
    ContextTransferTime  float64 `gorm:"type:decimal(16,6);default:0"`//:= t7.Sub(t4)
    TotalTime float64 `gorm:"type:decimal(16,6);default:0"` //:= t7.Sub(t0)
	StatusCode int16
	TextRegexMatch bool
    IPAddress  pgtype.Inet    `gorm:"type:inet"`
}

type Host struct {
//  IP  pgtype.Inet    `gorm:"type:inet"`
//   MAC pgtype.Macaddr `gorm:"type:macaddr"`
}

type Node struct {
}

type DomainName struct {
}



type URL struct {
	gorm.Model
	//URL_id       int64  `gorm:"primaryKey"`
	Disabled         bool `gorm:"type:bool;default:false"`
	UUID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Note             string
    HTTPConnections []HTTPConnection
    HTTPConnectionsTriggers []HTTPConnectionTrigger `gorm:"many2many:URLs_x_HTTPConnectionTriggers;"`
    Groups []Group `gorm:"many2many:URLs_x_Groups;"`
	Path             string `gorm:"unique;not null"`
	AllowInsecureTLS bool `gorm:"default:false`
	Test     string
	Test2    string
    Platforms []Platform `gorm:"many2many:url_x_platform;"`
	// Primary_Platform
	// Rsp_time     float64
}




// var URLS []URL{}
// var URLS []string
// func GetURLs() []string{
func GetURLs() []URL {
	URLS := []URL{}
	// URLS := []URL{}
	// URLS := []string{"a", "b", "c", "d"}
	//URLS := []URL
	// URLS := database.DBCon(&URL)
	// database.DBCon.First(&URL, "path = ?", '*')
	// database.DBCon.First(&URL, "path = ?", "https://example.com")
	database.DBCon.Find(&URLS)
	fmt.Printf("URLS: %v\n", URLS)
	for i, URL := range URLS {
		fmt.Printf("index: %v, value: %v\n", i, URL)
	}
	//fmt.Printf("URLS: %+v\n", URLS)
	return URLS
}

func CheckURLs() []URL {
	URLS := []URL{}
	database.DBCon.Find(&URLS)
	// fmt.Printf("URLS: %v\n", URLS)
	for i, URL := range URLS {
		fmt.Printf("index: %v, value: %v\n", i, URL)
		fmt.Printf("%i URL.Path: %v\n", i, URL.Path)
		getURLStats(URL.Path)
	}

	return URLS
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

func GetOneURL() URL {
	URL := URL{}
	// URLS := []URL{}
	// URLS := []string{"a", "b", "c", "d"}
	//URLS := []URL
	// URLS := database.DBCon(&URL)
	// database.DBCon.First(&URL, "path = ?", '*')
	database.DBCon.First(&URL, "path = ?", "https://example.com")
	fmt.Printf("URLS: %v\n", URL)
	return URL
}

func GetURL() {
	path := "https://example.com"
	URL := URL{}
	// URL = "{Path: Foo}"
	database.DBCon.First(&URL, "path = ?", path)

	// URL := models.URL{}
	database.DBCon.First(&URL, "path = ?", path)
	fmt.Printf("zzzz===============\n\n\n")
	database.DBCon.Model(&URL).Update("Rsp_time", 2)
	fmt.Printf("jjzzzz===============\n\n\n")
	fmt.Printf("URL Path: %v\n", URL.Path)
	fmt.Printf("URL UUID: %v\n", URL.UUID)
	// fmt.Printf("URL.path: %v", &URL)
	fmt.Printf("zzzz===============\n\n\n")
	// fmt.Printf("URL.path: %v\n", &URL.Path)
	fmt.Printf("URL.path: %v\n", &URL.UUID)
	os.Exit(0)
	fmt.Printf("zzzz===============\n\n\n")
	fmt.Printf("URL: %+v\n", &URL)
	fmt.Printf("YYY===============\n\n\n")

	fmt.Printf("aaaaaa===============\n")
	fmt.Printf("URL: %+v\n", &URL.ID)
	fmt.Printf("URL: %+v\n", &URL.ID)
	fmt.Printf("sssss===============\n")
	fmt.Printf("URL: %+v\n", &URL.ID)
	fmt.Printf("sssss===============\n")
	fmt.Printf("FF: %v\n", &URL.ID)
	fmt.Printf("bbbbbb===============\n")
	// database.DBCon.Query("hello")
}

// fmt.Println("FOO:", os.Getenv("FOO"))
// dbURL := os.Getenv("MINIMON_DBURL")
// DB := db.Init(dbURL)

// func init(){
// database.DBCon.AutoMigrate(
//        &User{},
//        &URL{} )
