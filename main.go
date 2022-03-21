package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	//"context"

	"github.com/go-co-op/gocron"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
	"minimon/db"
	"minimon/models"
	// "github.com/jackc/pgtype"
	// "github.com/jackc/pgx/v4"
	"context"
	// "reflect"
	// "minimon/monitorhttp"
)

//func AutoMigrate() {
//	err := db.DBCon.AutoMigrate(
//		&models.Platform{},
//		&models.HTTPConnectionTrigger{},
//		&models.URL{},
//		&models.User{})
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func main() {
fmt.Printf("func main\n")
	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		// do something with every new connection
	// }
	// var err error
	db.DB, err = pgxpool.ConnectConfig(context.Background(), config)
	// testPgx()
	testPgxpool()
	testPgxpool()

	// models.AutoMigrate()
	// models.AutoMigrate()

	// testing()
	//models.GetURLs()
	// models.CheckURLs()
	//var r models.HTTPConnection
	// r := monitorhttp.HTTPRequest("https://www.uvoo.io")
	//fmt.Printf("Return: %v\n", r)
	// 	dotask()
	// serve()

	// monitor()

}

// ===============NOTES ========================================================

func monitor() {
	// models.GetURLs()
	// fmt.Printf("foo\n")
	models.CheckURLs()
	//jfmt.Printf("aaa: %+v\n", a)
	// fmt.Printf("foo\n")
}

func task() {
	fmt.Println("I am running task.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func testPgx() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	// conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var path string
	var id int64
	err = conn.QueryRow(context.Background(), "select id, path from urls where id=$1", 1).Scan(&id, &path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(id, path)
}

// func testPgxpool(dbPool *pgxpool.Pool) {
func testPgxpool() {
fmt.Printf("func tesPgxpool\n")
	// execute the select query and get result rows
	// dbPool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	// to close DB pool
	// defer dbPool.Close()
	// rows, err := dbPool.Query(context.Background(), "select id, path from urls")
	rows, err := db.DB.Query(context.Background(), "select id, path from urls")
	if err != nil {
		log.Fatal("error while executing query")
	}

	// iterate through the rows
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}

		// convert DB types to Go types
		id := values[0].(int32)
		path := values[1].(string)
		// lastName := values[2].(string)
		// dateOfBirth := values[3].(time.Time)
		// log.Println("[id:", id, ", first_name:", firstName, ", last_name:", lastName, ", date_of_birth:", dateOfBirth, "]")
		log.Println("[id:", id, ", path:", path, "]")
	}

}

func dotask() {

	// defines a new scheduler that schedules and runs jobs
	s1 := gocron.NewScheduler(time.UTC)

	s1.Every(3).Seconds().Do(task)

	// scheduler starts running jobs and current thread continues to execute
	s1.StartAsync()
	s1.StartBlocking()

	// s := gocron.NewScheduler(time.UTC)

	// s.Every(5).Seconds().Do(func(){ fmt.Println("Hi & Hello") })
	//  s.Every(5).Seconds().Do(func(){ log.Println("Hello world!") })
	// s.Every(5).Seconds().Do( hi() )
}

func hi() {
	// fmt.Println("Hi & Hello")
	log.Println("Hello world!")
}

func serve() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.POST("/urls", saveUser)
	e.GET("/urls/:id", getUser)
	e.PUT("/urls/:id", updateUser)
	e.DELETE("/urls/:id", deleteUser)

	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

//e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// e.POST("/save", save)
func saveUser(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

func updateUser(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}
func deleteUser(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

func testing() {
	path := "https://example.com"
	URL := models.URL{}
	db.DBCon.First(&URL, "path = ?", path)
	fmt.Printf("URL Path: %v\n", URL.Path)
	fmt.Printf("URL ID : %v\n", URL.ID)

	fmt.Printf("===========")
	// foo.Get()
	models.GetURL()
	// fmt.Printf("%v", a)
	fmt.Printf("===========")
	// e := foo.run()
}
