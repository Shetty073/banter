package migrations

import (
	"banter/models"
	"banter/stores"
	"banter/utils/logger"
)

func RegisterAllModels() {
	modelTypes := []interface{}{
		&models.User{},
		&models.Role{},
		&models.Module{},
		&models.Permission{},
		&models.RoleModulePermission{},
		&models.UserRole{},

		&models.Address{},
		&models.ServicablePinCode{},
		&models.Category{},
		&models.SubCategory{},
		&models.Warehouse{},
		&models.Currency{},

		&models.Brand{},
		&models.Product{},
		&models.ProductMedia{},
		&models.Inventory{},
		&models.InventoryProduct{},
		&models.Supplier{},

		&models.PurchaseOrder{},
		&models.SalesOrder{},
		&models.Logistics{},
		&models.LogisticsUpdate{},
		&models.InventoryTransaction{},
		&models.MarketingCampaign{},

		&models.Promotion{},
		&models.SupportTicket{},
		&models.TicketNote{},
		&models.Chat{},
		&models.ChatMessage{},
	}

	// Iterate through all models registered in the AllModels slice and auto-migrate them
	for _, model := range modelTypes {
		if err := stores.GetDb().AutoMigrate(model); err != nil {
			logger.Logger.Fatalf("Failed to auto-migrate model %v: %v", model, err)
		}
	}
}
