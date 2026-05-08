package main

import "fmt"

type MyOrder string

const (
	Recieved  MyOrder = "recieved"
	Preparing MyOrder = "preparing"
	Delivered MyOrder = "delivered"
)

func changeOrderStatus(status MyOrder) {
	fmt.Println("You order status: ", status)
}

func main() {
	changeOrderStatus(Preparing)
}
