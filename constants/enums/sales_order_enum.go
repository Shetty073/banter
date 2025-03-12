package enums

type SalesOrderStatus string

const (
	SalesOrderPending   SalesOrderStatus = "pending"
	SalesOrderCompleted SalesOrderStatus = "shipped"
	SalesOrderCancelled SalesOrderStatus = "delivered"
)
