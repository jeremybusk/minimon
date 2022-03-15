package models

import (
	// "database/sql"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"minimon/database"
	"os"
	"github.com/tcnksm/go-httpstat"
	"log"
	"net/http"
	"time"
	"io"
	"io/ioutil"

	// https://github.com/davecheney/httpstat.git
	// "time"
)

type URL struct {
	gorm.Model
	//URL_id       int64  `gorm:"primaryKey"`
	Disabled         bool
	UUID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Note             string
	URL_group_id     int
	Path             string `gorm:"unique;not null"`
	Rsp_code         int
	Rsp_code_exp     int `gorm:"default:200`
	Rsp_code_test    bool
	Rsp_time         float64 `gorm:"type:decimal(16,6);default:0"`
	Rsp_time_exp     int     `gorm:"default:4`
	Rsp_time_test    bool
	Rsp_regex_exp    string `gorm:"default:statushealthy`
	Rsp_regex_test   bool
	AllowInsecureTLS bool `gorm:"default:false`
// exp is expected or threshold value
	DNS_lookup_rsp_time float64 `gorm:"type:decimal(16,6);default:0"`
	DNS_lookup_rsp_time_exp float64 `gorm:"type:decimal(16,6);default:0"`
	DNS_lookup_rsp_time_test bool `gorm:"type:int;default:0"`
	TCP_connection_rsp_time float64 `gorm:"type:decimal(16,6);default:0"`
	TCP_connection_rsp_time_exp float64 `gorm:"type:decimal(16,6);default:0"`
	TCP_connection_rsp_time_test bool `gorm:"type:int;default:0"`
	TLS_handshake_rsp_time float64 `gorm:"type:decimal(16,6);default:0"`
	TLS_handshake_rsp_time_exp float64 `gorm:"type:decimal(16,6);default:0"`
	TLS_handshake_rsp_time_test bool `gorm:"type:int;default:0"` 
	Server_processing_rsp_time float64 `gorm:"type:decimal(16,6);default:0"`
	Server_processing_rsp_time_exp float64 `gorm:"type:decimal(16,6);default:0"`
	Server_processing_rsp_time_test bool `gorm:"type:int;default:0"`
	Content_transfer_rsp_time float64 `gorm:"type:decimal(16,6);default:0"`
	Content_transfer_rsp_time_exp float64 `gorm:"type:decimal(16,6);default:0"`
	Content_transfer_rsp_time_test bool `gorm:"type:int;default:0"`

	// Rsp_time     Decimal `gorm:"type:decimal(16,6);default:0"`
	//Amount       float32   `sql:"type:decimal(10,2);"`
	Sequence int
	Test     string
	Test2    string
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
