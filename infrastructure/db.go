package infrastructure

import (
	"fmt"
	"rnd-surajan-gin/environment"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

// This database connection function will run automatically once put in "fx.Options" in "bootstrap.go" file after exporting as a "Module" in "infrastructure.go".
func NewDatabase() Database {
	// DB connection
	dbCred := environment.GetDatabaseEnv()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbCred[0], dbCred[1], dbCred[2], dbCred[3], dbCred[4])
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return Database{DB: db}
}
