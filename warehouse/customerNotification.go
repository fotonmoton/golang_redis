package warehouse

// Notifies a customer about product for which customer has subscription:
// 1, 1 (Greg, "Blue Jeans")
// 2, 2 (Bob, "Gray T-shirt")
type CustomerNotification struct {
	CustomerId int
	ProductId  int
}
