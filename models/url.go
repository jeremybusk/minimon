package models

import (
	// "db/sql"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/tcnksm/go-httpstat"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"log"
	"minimon/db"
	"minimon/monitorhttp"
	"net"
	"net/http"
	"os"
	"time"
	//"github.com/jackc/pgtype"
	// https://github.com/davecheney/httpstat.git
	// "time"
)

// sql.NullString
// Int64 32 16 Time
type Platform struct {
	gorm.Model
	Disabled bool      `gorm:"type:bool;default:false"`
	UUID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Note     sql.NullString
	Sequence int `gorm:"type:int;default:0"`
	Name     sql.NullString
}

type Group struct {
	gorm.Model
	//URL_id       int64  `gorm:"primaryKey"`
	Disabled bool      `gorm:"type:bool;default:false"`
	UUID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Note     sql.NullString
	name     string
}

//fmt.Println(namelookup, connect, pretransfer, starttransfer, total)
type HTTPConnectionTrigger struct {
	gorm.Model
	UUID                 uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4()"`
	DomainNameLookupTime time.Duration `gorm:"type:decimal(16,6);default:500"`   //:= t1.Sub(t0)
	TCPConnectionTime    time.Duration `gorm:"type:decimal(16,6);default:500"`   //:= t2.Sub(t1)
	ConnectTime          time.Duration `gorm:"type:decimal(16,6);default:500"`   //:= t2.Sub(t0)
	PreTransferTime      time.Duration `gorm:"type:decimal(16,6);default:500"`   //:= t3.Sub(t0)
	StartTransferTime    time.Duration `gorm:"type:decimal(16,6);default:500"`   //:= t4.Sub(t0)
	ServerProcessingTime time.Duration `gorm:"type:decimal(16,6);default:500"`   //:= t4.Sub(t3)
	TLSHandshakeTime     time.Duration `gorm:"type:decimal(16,6);default:500"`   //:= t6.Sub(t5)
	ContextTransferTime  time.Duration `gorm:"type:decimal(16,6);default:500"`   //:= t7.Sub(t4)
	TotalTime            time.Duration `gorm:"type:decimal(16,6);default:10000"` //:= t7.Sub(t0)
	StatusCode           int           `gorm:"default:200`
	CheckResponseBody    bool          `gorm:"type:bool;default:false"`
	CheckResponseHeader  bool          `gorm:"type:bool;default:false"`
	ResponseBodyRegex    string        `gorm:"default:megamonstatushealthy`
	ResponseHeaderRegex  string        `gorm:"default:megamonstatushealthy`
	// IPAddress  pgtype.Inet    `gorm:"type:inet;default:'0.0.0.0/0'"`
	IPAddress net.IP `gorm:"type:inet;default:'0.0.0.0/0'"`
	Note      string
	Test1     string
	URLs      []URL
}

//fmt.Println(namelookup, connect, pretransfer, starttransfer, total)
type HTTPConnection struct {
	gorm.Model
	UUID                 uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	URLID                int
	DomainNameLookupTime time.Duration //:= t1.Sub(t0)
	TCPConnectionTime    time.Duration //:= t2.Sub(t1)
	ConnectTime          time.Duration //:= t2.Sub(t0)
	PreTransferTime      time.Duration //:= t3.Sub(t0)
	StartTransferTime    time.Duration //:= t4.Sub(t0)
	ServerProcessingTime time.Duration //:= t4.Sub(t3)
	TLSHandshakeTime     time.Duration //:= t6.Sub(t5)
	ContextTransferTime  time.Duration //:= t7.Sub(t4)
	TotalTime            time.Duration //:= t7.Sub(t0)
	StartTime            time.Time     //:= t0
	StopTime             time.Time     //:= t7
	StatusCode           int
	//TestResponseBodyRegex int
	//TestResponseHeaderRegex int
	// TextRegexMatch       bool
	// IPAddress  pgtype.Inet    `gorm:"type:inet"`
	IPAddress                net.IP `gorm:"type:inet"`
	ResponseBodyRegexMatch   bool
	ResponseHeaderRegexMatch bool
	ResponseBodyText         string
	ResponseHeaderText       string
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
	Disabled        bool      `gorm:"type:bool;default:false"`
	UUID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Note            string
	HTTPConnections []HTTPConnection
	// HTTPConnectionsTriggers []HTTPConnectionTrigger `gorm:"many2many:URLs_x_HTTPConnectionTriggers;"`
	Groups                  []Group    `gorm:"many2many:URLs_x_Groups;"`
	Path                    string     `gorm:"unique;not null"`
	AllowInsecureTLS        bool       `gorm:"default:false`
	Platforms               []Platform `gorm:"many2many:url_x_platform;"`
	HTTPConnectionTriggerID uint       `gorm:"default:1"`
	// Primary_Platform
	// Rsp_time     float64
}

// var URLS []URL{}
// var URLS []string
// func GetURLs() []string{
func GetURLs() []URL {
	fmt.Printf("===============\n")
	URLS := []URL{}
	// URLS := []URL{}
	// URLS := []string{"a", "b", "c", "d"}
	//URLS := []URL
	// URLS := db.DBCon(&URL)
	// db.DBCon.First(&URL, "path = ?", '*')
	// db.DBCon.First(&URL, "path = ?", "https://example.com")
	db.DBCon.Find(&URLS)
	fmt.Printf("URLS: %+v\n", URLS)
	for i, URL := range URLS {
		fmt.Printf("index: %v, value: %v\n", i, URL)
	}
	//fmt.Printf("URLS: %+v\n", URLS)
	return URLS
}

func CheckURLs() []URL {
	fmt.Printf("===============\n")
	URLS := []URL{}
	HTTPConnectionTriggers := []HTTPConnectionTrigger{}
	// HTTPConnectionTriggers := []HTTPConnectionTrigger{}
	db.DBCon.Find(&URLS)
	// db.DBCon.Preload("HTTPConnectionTriggers").Find(&URLS)
	// fmt.Printf("URLS: %v\n", URLS)
	for i, URL := range URLS {
		fmt.Printf("index: %v, value: %+v\n", i, URL)
		db.DBCon.Find(&HTTPConnectionTriggers, URL.ID)
		fmt.Printf("===============\n\n")
		fmt.Printf("&HTTPConnectionTrigger %+v\n", HTTPConnectionTriggers)
		fmt.Printf("===============\n\n")
		//db.DBCon.Find(&HTTPConnectionTriggers)
		//db.DBCon(&HTTPConnectionTriggers).First(&result, "id = ?", URL.id)
		//db.DBCon.Preload("HTTPConnectionTriggers").Find(&services)
		//db.Take(&user).Where()

		// fmt.Printf("%i URL.Path: %v\n", i, URL.Path)
		// r := monitorhttp.HTTPRequest("https://uvoo.io.com")
		// fmt.Printf("====================")
		// r := monitorhttp.HTTPRequest(URL.Path, HTTPConnectionTrigger.ResponseBodyRegex, HTTPConnectionTrigger.ResponseHeaderRegex)
		// r := monitorhttp.HTTPRequest(URL.Path, URLHTTPConnectionTrigger)

		r := monitorhttp.HTTPRequest(URL.Path)
		fmt.Printf("%v %v\n", i, r.TotalTime)

		// fmt.Printf("====================%v\n", r)
		// getURLStats(URL.Path)
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
	// URLS := db.DBCon(&URL)
	// db.DBCon.First(&URL, "path = ?", '*')
	db.DBCon.First(&URL, "path = ?", "https://example.com")
	fmt.Printf("URLS: %v\n", URL)
	return URL
}

func GetURL() {
	path := "https://example.com"
	URL := URL{}
	// URL = "{Path: Foo}"
	db.DBCon.First(&URL, "path = ?", path)

	// URL := models.URL{}
	db.DBCon.First(&URL, "path = ?", path)
	fmt.Printf("zzzz===============\n\n\n")
	db.DBCon.Model(&URL).Update("Rsp_time", 2)
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
	// db.DBCon.Query("hello")
}

// fmt.Println("FOO:", os.Getenv("FOO"))
// dbURL := os.Getenv("MINIMON_DBURL")
// DB := db.Init(dbURL)

// func init(){
// db.DBCon.AutoMigrate(
//        &User{},
//        &URL{} )
