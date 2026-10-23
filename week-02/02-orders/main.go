package main

import "fmt"

func main() {
	orders := []int{12000, 7500, 31000, 4500}
	if len(orders) == 0 {
		fmt.Println("No orders to analyze")
		return
	}

	minAmount := orders[0]
	maxAmount := orders[0]
	var total int

	for _, amount := range orders {
		total += amount
		if amount < minAmount {
			minAmount = amount
		}
		if amount > maxAmount {
			maxAmount = amount
		}
	}

	average := float64(total) / float64(len(orders))

	fmt.Println("Order count:", len(orders))
	fmt.Println("Minimum order:", minAmount, "kopecks")
	fmt.Println("Maximum order:", maxAmount, "kopecks")
	fmt.Println("Total amount:", total, "kopecks")
	fmt.Println("Average order:", average, "kopecks")
}
