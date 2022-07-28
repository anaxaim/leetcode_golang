package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	min, maxSale := prices[0], 0

	for _, price := range prices {
		if price < min {
			min = price
		} else if (price - min) > maxSale {
			maxSale = price - min
		}
	}

	return maxSale
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
