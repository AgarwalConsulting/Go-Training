package main

type Customer struct {
	Name         string
	WalletAmount float64
}

type MenuItem struct {
	Name string
}

type RestaurantMenu struct {
	MenuItem MenuItem
	Price    float64
}

type Restaurant struct {
	Name            string
	RestaurantMenus []RestaurantMenu
}

type Order struct {
	OrderNumber  string
	Status       string
	Restaurant   Restaurant
	Customer     Customer
	OrderDetails []OrderDetail
}

type OrderDetail struct {
	Order      Order
	ItemName   string
	Quantity   uint
	Price      float64
	TotalPrice float64
}

type OrderService struct{}

func (os OrderService) PlaceOrder(order Order) Order {}

func (os OrderService) Confirmation(order Order) bool {
	return PaymentService{}.Pay(order)
}

type PaymentService struct{}

func (ps PaymentService) Pay(order Order) {}
