package infrastructure

import "rnd-surajan-gin/models"

type Migrations struct {
	db Database
}

func NewMigrations(db Database) Migrations {
	return Migrations{db: db}
}

func (cc Migrations) Migrate() {
	// Migrate models.
	cc.db.DB.AutoMigrate(&models.Task{})
	cc.db.DB.AutoMigrate(&models.User{})
}
