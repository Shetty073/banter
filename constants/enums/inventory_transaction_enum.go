package enums

type InventoryTransactionStatus string

const (
	InventoryTransactionIn         InventoryTransactionStatus = "in"
	InventoryTransactionOut        InventoryTransactionStatus = "out"
	InventoryTransactionAdjustment InventoryTransactionStatus = "adjustment"
)
