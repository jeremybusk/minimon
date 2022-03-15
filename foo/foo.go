package foo

import "myapp/database"
import "myapp/models"
import "fmt"
func Get() {
path := "https://example.com"
URL := models.URL{}
// URL = "{Path: Foo}"
    database.DBCon.First(&URL, "path = ?", path)
	// fmt.Printf("URL.path: %v", &URL)
	fmt.Printf("zzzz===============\n\n\n")
	fmt.Printf("URL.path: %v\n", &URL.Path)
	fmt.Printf("zzzz===============\n\n\n")
	fmt.Printf("URL: %+v\n", &URL)
	fmt.Printf("YYY===============\n\n\n")

	fmt.Printf("aaaaaa===============\n")
	fmt.Printf("URL: %+v\n", &URL.Model.ID)
	fmt.Printf("URL: %+v\n", &URL.Model.ID)
	fmt.Printf("sssss===============\n")
	fmt.Printf("URL: %+v\n", &URL.Model.ID)
	fmt.Printf("sssss===============\n")
	fmt.Printf("FF: %v\n", &URL.Model.ID)
	fmt.Printf("bbbbbb===============\n")
   // database.DBCon.Query("hello")
}
