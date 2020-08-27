package main

import (
	"github.com/omkarjoshi/go-json-examples/objects"
)

func main() {
	order := objects.Order{
		OrderId:         "OR1D",
		TotalItemsCount: 0,
	}
	println(order.OrderId)

}
