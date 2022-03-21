package models

import (
"minimon/database"
)

func migrate() {
    database.DBCon.AutoMigrate(
	&URL{},
	&User{})
}
