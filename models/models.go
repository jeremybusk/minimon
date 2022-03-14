package models

import(
"gorm.io/gorm"
"github.com/google/uuid"
// "gorm.io/driver/postgres"
"time"
"database/sql"
"github.com/uvoo/minimon/pkg/db"
"os"
//"github.com/uvoo/minimon/pkg/db"
)

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
    AllowInsecureTLS    bool `gorm:"default:false`
    // Rsp_time     Decimal `gorm:"type:decimal(16,6);default:0"`
    //Amount       float32   `sql:"type:decimal(10,2);"`
    Sequence     int
	Test string
	Test2 string
    // Rsp_time     float64
}


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

// fmt.Println("FOO:", os.Getenv("FOO"))
func init(){
dbURL := os.Getenv("MINIMON_DBURL")
DB := db.Init(dbURL)
DB.AutoMigrate(
	&User{},
	&URL{} )
}
