package models

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"minimon/database"
	// "os"
	"time"
)

type User struct {
	gorm.Model
	UUID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
    Disabled         bool
    Note             string
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

func Ghet() {
	path := "https://example.com"
	URL := URL{}
	database.DBCon.First(&URL, "path = ?", path)

	fmt.Printf("URL Path: %v\n", URL.Path)
	fmt.Printf("URL UUID: %v\n", URL.UUID)
}
