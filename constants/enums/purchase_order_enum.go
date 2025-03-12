package enums

type PurchaseOrderStatus string

const (
	PurchaseOrderPending   PurchaseOrderStatus = "pending"
	PurchaseOrderCompleted PurchaseOrderStatus = "completed"
	PurchaseOrderCancelled PurchaseOrderStatus = "cancelled"
)
