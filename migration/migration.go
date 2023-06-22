package migration

import (
	"rnd-surajan-gin/database"
	"rnd-surajan-gin/models"
)

func SyncDatabase() {
	// Migrate the schema
	database.DB.AutoMigrate(&models.Task{})
}
