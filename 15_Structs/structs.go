package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone int
}

// order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
	customer            // struct embedding
}

//constructor (custom made)

func newOrder(id string, amount float32, status string) *order {
	newOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &newOrder
}

func (o *order) changeStatus(status string) {
	o.status = status
}

func main() {
	// var order order

	// order := order{
	// 	id:     "1",
	// 	amount: 50.00,
	// 	status: "dispatched",
	// }
	// order.createdAt = time.Now()
	// fmt.Println(order)
	// order.changeStatus("delivered")
	// fmt.Println(order)

	myOrder := newOrder("1", 50.50, "Shipped")
	fmt.Println(myOrder)

	//inline struct
	language := struct {
		name   string
		isGood bool
	}{"GoLang", true}
	fmt.Println(language)

	emOrder := order{
		id:     "2",
		amount: 80.55,
		status: "pending",
		customer: customer{
			name:  "John",
			phone: 1234567890,
		},
	}
	fmt.Println(emOrder)
	emOrder.customer.name = "John Doe"
	fmt.Println(emOrder)
}
