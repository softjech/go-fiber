package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Define variable to connect to the database
var (
	DBConn *gorm.DB
)
