package database

import (
	"gorm.io/gorm"
)

var (
	// DBCon is the connection handle database
	DBCon *gorm.DB
)

// Usage:
// URL := models.URL{}
// database.DBCon.First(&URL, "path = ?", "my string")
// fmt.Printf("URL Path: %v\n", URL.Path)
