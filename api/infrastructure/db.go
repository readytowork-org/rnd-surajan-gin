package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

// This database connection function will run automatically once put in "fx.Options" in "bootstrap.go" file after exporting as a "Module" in "infrastructure.go".
func NewDatabase() Database {
	// DB connection
	dsn := "root:password@123@tcp(localhost:3306)/rnd-gin?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return Database{DB: db}
}
