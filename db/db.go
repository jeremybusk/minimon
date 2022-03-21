package db

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"gorm.io/gorm"
)

var (
	// DBCon is the connection handle database
	DBCon *gorm.DB
)

var (
	// DBCon is the connection handle database
	DB *pgxpool.Pool
)

// Usage:
// URL := models.URL{}
// database.DBCon.First(&URL, "path = ?", "my string")
// fmt.Printf("URL Path: %v\n", URL.Path)
