package main

import "fmt"

func main() {
	amount := 25000
	if amount < 0 {
		fmt.Println("Error: order amount cannot be negative")
		return
	}

	var category string
	switch {
	case amount < 10000:
		category = "small"
	case amount < 50000:
		category = "regular"
	default:
		category = "large"
	}

	fmt.Println("Order amount:", amount, "kopecks")
	fmt.Println("Category:", category)
}
