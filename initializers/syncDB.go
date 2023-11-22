package initializers

import (
	"example.com/m/models"
)

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
