# Task Description

We have a food delivery application that connects between a customer, a restaurant.
The customer can order any food, and the system will send the order to the restaurant , then the system will summarize the order and send the result to the customer.
When the order is confirmed, the Order will be paid using user wallet

The app currently contains:

Entities:

- Customer
- MenuItem
- Restaurant
- RestaurantMenu
- Order
- OrderDetail

Services :

- OrderService
- PaymentService

Skeleton:

```golang
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
```

## Problem Statement

- System should support multiple payment methods (wallet, cash and credit card)
  - We can choose any payment we want to use
    - Also support for split payment using available payment methods

- Every weekend, every customer will get a 10% discount.
  - Customer will get additional 5% if the age more than 50 years old
  - Customer will get additional 5% if they order on their birthday

## Todo

- Design the solution for problem statements above
- Use best practice approach to implement the solution (TDD, Design Pattern, etc)
