package infrastructure

import "rnd-surajan-gin/models"

type Migrations struct {
	db Database
}

func NewMigrations(db Database) Migrations {
	return Migrations{db: db}
}

func (cc Migrations) Migrate() {
	// Migrate "Task" model.
	cc.db.DB.AutoMigrate(&models.Task{})
}
