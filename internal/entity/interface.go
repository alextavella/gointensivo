package entity

type OrderRepositoryInterface interface {
	SaveOrder(order *Order) error
	GetTotal() (int, error)
}
