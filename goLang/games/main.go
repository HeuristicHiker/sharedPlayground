package main

import (
	"fmt"
	"strconv"
)

func main() {

}

func calPoints(operations []string) int {
	sum := 0

	score := []int{}

	for idx, op := range operations {
		fmt.Println(idx, op)
		switch op {
		case "+":
			num1, num2 := score[len(score)-1], score[len(score)-2]
			score = append(score, num1+num2)
		case "C":
			score = score[:len(score)-1]
		case "D":
			fmt.Println("Length of stack: ", len(score), len(score)-1)
			score = append(score, score[len(score)-1]*2)
		default:
			num, _ := strconv.Atoi(op)
			score = append(score, num)
		}
	}

	for _, num := range score {
		sum += num
	}
	return sum
}

/*
There is a robot starting at the position (0, 0), the origin, on a 2D plane. Given a sequence of its moves, judge if this robot ends up at (0, 0) after it completes its moves.

You are given a string moves that represents the move sequence of the robot where moves[i] represents its ith move. Valid moves are 'R' (right), 'L' (left), 'U' (up), and 'D' (down).

Return true if the robot returns to the origin after it finishes all of its moves, or false otherwise.

Note: The way that the robot is "facing" is irrelevant. 'R' will always make the robot move to the right once, 'L' will always make it move left, etc. Also, assume that the magnitude of the robot's movement is the same for each move.



Example 1:

Input: moves = "UD"
Output: true
Explanation: The robot moves up once, and then down once. All moves have the same magnitude, so it ended up at the origin where it started. Therefore, we return true.
Example 2:

Input: moves = "LL"
Output: false
Explanation: The robot moves left twice. It ends up two "moves" to the left of the origin. We return false because it is not at the origin at the end of its moves.


Constraints:

1 <= moves.length <= 2 * 104
moves only contains the characters 'U', 'D', 'L' and 'R'.
*/

func judgeCircle(moves string) bool {
	vDiff := 0
	hDiff := 0

	for _, s := range moves {
		fmt.Println(s)
		switch s {
		case 'U':
			vDiff++
		case 'D':
			vDiff--
		case 'L':
			hDiff--
		case 'R':
			hDiff++
		}
	}

	return hDiff == 0 && vDiff == 0
}
