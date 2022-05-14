package main

import (
	"context"
	"fmt"
	"stepik_golang/grpcteta/pkg/orderservice"
	"time"

	"google.golang.org/grpc"
)

func main() {

	cwt, _ := context.WithTimeout(context.Background(), time.Second*5)
	conn, err := grpc.DialContext(cwt, "localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	os := orderservice.NewOrdersClient(conn)

	newOrder := &orderservice.Order{
		Name:     "myphone",
		Price:    "10000",
		Category: "smartphone",
	}

	orderResponse, err := os.CreateOrder(cwt, newOrder)
	if err != nil {
		panic(err)
	}

	if orderResponse.Success {
		fmt.Println("Order created")
	}

	orders, err := os.GetOrders(cwt, &orderservice.Filter{})
	if err != nil {
		panic(err)
	}
	fmt.Println(orders)
}
