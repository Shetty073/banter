package migrations

import (
	"banter/models"
	"banter/stores"
	"banter/utils/logger"
)

func RegisterAllModels() {
	modelTypes := []interface{}{
		&models.User{},
		// add new models here for migration
	}

	// Iterate through all models registered in the AllModels slice and auto-migrate them
	for _, model := range modelTypes {
		if err := stores.GetDb().AutoMigrate(model); err != nil {
			logger.Logger.Fatalf("Failed to auto-migrate model %v: %v", model, err)
		}
	}
}
