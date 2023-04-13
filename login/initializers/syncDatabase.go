package initializers

import "login/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}