package orders

type Order struct {
	ID         string
	Status     string
	TotalCents int
}

const (
	StatusDelivered = "delivered"
	StatusCanceled  = "canceled"
	StatusPending   = "pending"
)

func CanReturn(order Order) bool {
	return order.Status == StatusDelivered
}
