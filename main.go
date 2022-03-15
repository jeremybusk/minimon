package main

import (
	"fmt"
	"log"
  "net/http"
	"os"
	"time"
	//"context"

  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/go-co-op/gocron"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"minimon/database"
	"minimon/models"
)

func main() {
	var err error
	dbURL := os.Getenv("MINIMON_DBURL")
	database.DBCon, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// testing()
  //models.GetURLs()
  models.CheckURLs()
// 	dotask()
	// serve()


}



// ===============NOTES ========================================================

func monitor(){
  models.GetURLs()
}

func task() {
    fmt.Println("I am running task.")
}

func taskWithParams(a int, b string) {
    fmt.Println(a, b)
}

func dotask(){

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

func hi(){
// fmt.Println("Hi & Hello")
log.Println("Hello world!")
}


func serve(){
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
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}


// e.POST("/save", save)
func saveUser(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}

func updateUser(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}
func deleteUser(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}



func testing(){
	path := "https://example.com"
	URL := models.URL{}
	database.DBCon.First(&URL, "path = ?", path)
	fmt.Printf("URL Path: %v\n", URL.Path)
	fmt.Printf("URL ID : %v\n", URL.ID)

	fmt.Printf("===========")
	// foo.Get()
	models.GetURL()
	// fmt.Printf("%v", a)
	fmt.Printf("===========")
	// e := foo.run()
}
