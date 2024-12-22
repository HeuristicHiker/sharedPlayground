package main

import "fmt"

func main() {
	
}

func diagonalSum(matrix [][]int) int {
	length := len(matrix[0]) - 1
	sum := 0

	for i := 0; i <= length; i++ {
		dE := matrix[i][i]
		aE := matrix[i][length-i]
		sum += dE + aE
		fmt.Println(dE, aE)
	}

	if (length+1)%2 != 0 {
		sum -= matrix[length/2][length/2]
	}

	return sum
}
