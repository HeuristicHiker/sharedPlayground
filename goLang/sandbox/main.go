package main

import (
	"fmt"
)

func main() {
	testCase1 := [][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}
	fmt.Println(tictactoe(testCase1, 3))
}

func printBoard(board [][]int) {
	for _, row := range board {
		fmt.Println(row)
	}
}

func checkSums(board [][]int) string {
	n := len(board)
	for idx := 0; idx < n; idx++ {
		horizontal := 0
		vertical := 0
		diagonal := 0
		inverseDiagonal := 0
		for i := 0; i < n; i++ {
			horizontal += board[idx][i]
			vertical += board[i][idx]
			diagonal += board[i][i]
			inverseDiagonal += board[i][n-1-i]
		}
		if horizontal == n || vertical == n || diagonal == n || inverseDiagonal == n {
			return "A"
		} else if horizontal == -n || vertical == -n || diagonal == -n || inverseDiagonal == -n {
			return "B"
		}
	}
	return "Pending"
}

func tictactoe(moves [][]int, n int) string {
	boardState := make([][]int, n)
	for i := range boardState {
		boardState[i] = make([]int, n)
	}
	isA := true

	for _, rc := range moves {
		row := rc[0]
		col := rc[1]
		if isA {
			boardState[row][col] = 1
			isA = false
		} else {
			boardState[row][col] = -1
			isA = true
		}
		result := checkSums(boardState)
		if result != "Pending" {
			return result
		}
	}

	if len(moves) >= n*n {
		return "Draw"
	}

	return "Pending"
}
