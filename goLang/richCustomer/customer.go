package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello from cust")
}

func richestCustomer(matrix [][]int) int {
	richestCust := 0
	mostMoney := 0

	for c, row := range matrix {
		totalM := 0
		for _, m := range row {
			totalM += m
		}
		if totalM > mostMoney {
			mostMoney = totalM
			richestCust = c
		}
	}

	return richestCust
}
