package initializers

import (
	"github.com/yosuahres/go-backend/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}