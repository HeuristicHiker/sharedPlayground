package main

import (
	"testing"

	"github.com/HeuristicHiker/shared-playground/goLang/ticTacToe"
)

func TestTicTacToe(t *testing.T) {
	tests := []struct {
		moves [][]int
		n     int
		want  string
	}{
		// Test case 1: Player A wins
		{
			moves: [][]int{{0, 0}, {1, 1}, {0, 1}, {1, 2}, {0, 2}},
			n:     3,
			want:  "A",
		},
		// Test case 2: Player B wins
		{
			moves: [][]int{{0, 0}, {1, 1}, {0, 1}, {1, 2}, {2, 0}, {2, 2}},
			n:     3,
			want:  "B",
		},
		// Test case 3: Draw
		{
			moves: [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 0}, {2, 2}},
			n:     3,
			want:  "Draw",
		},
		// Test case 4: Pending
		{
			moves: [][]int{{0, 0}, {1, 1}, {0, 1}, {1, 2}},
			n:     3,
			want:  "Pending",
		},
		// Test case 5: Larger board, Player A wins
		{
			moves: [][]int{{0, 0}, {1, 1}, {0, 1}, {1, 2}, {0, 2}, {2, 2}, {0, 3}},
			n:     4,
			want:  "A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := ticTacToe.Tictactoe(tt.moves, tt.n); got != tt.want {
				t.Errorf("Tictactoe(%v, %d) = %v; want %v", tt.moves, tt.n, got, tt.want)
			}
		})
	}
}
