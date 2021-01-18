package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var stepCounter int
var totalTimeTaken time.Duration

func main() {

	startTime := time.Now()

init:

	puzzle := [3][3]int{}
	puzzle = initBoard(puzzle)

	if !isBoardSolvable(puzzle) {
		goto init
	}

	boardTimeTaken := time.Since(startTime)
	fmt.Println("Time taken to generate the board:", boardTimeTaken)

	originalBoard := puzzle

	solveTimeStart := time.Now()

start:
	fmt.Print("\n\n")
	printBoard(puzzle)

	puzzle = solveBoard(puzzle)

	if !isBoardFinished(puzzle) {
		goto start
	}

	totalTimeTaken = time.Since(solveTimeStart)
	boardFinished(puzzle, originalBoard)
	fmt.Print("\n\n")
}

/* TODO- complete this function to solve the board */
func solveBoard(board [3][3]int) [3][3]int {

	newBoard := board

	positionGapI, positionGapJ := findGap(newBoard)
	fmt.Println("Gap is at:", positionGapI, positionGapJ)

	// TODO- steps to solve the board:
	//* bring 1 to it's position
	//* bring 3 to 2's place or 2 to 3's place
	//* bring 2 right below 3 or 3 right below 2
	//* rotate the pieces so that 2 and 3 are in the right position
	//* bring 7 to 4's place
	//* bring 4 to 5's place
	//* rotate anti-clockwise, so that 4 and 7 are in their right positions
	//* rotate the remaining 3 numbers, till the board is complete

	return newBoard
}

func isAlreadyPresent(board [3][3]int, x int) bool {

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == x {
				return true
			}
		}
	}
	return false
}

func getUniqueAndRandomNum(board [3][3]int) int {

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(9) + 1

	if !isAlreadyPresent(board, x) {
		return x
	}
	return getUniqueAndRandomNum(board)
}

func printBoard(board [3][3]int) {

	fmt.Println("| ", getNum(board, 0, 0), " | ", getNum(board, 0, 1), " | ", getNum(board, 0, 2), " |")
	fmt.Println("| ", getNum(board, 1, 0), " | ", getNum(board, 1, 1), " | ", getNum(board, 1, 2), " |")
	fmt.Println("| ", getNum(board, 2, 0), " | ", getNum(board, 2, 1), " | ", getNum(board, 2, 2), " |")
}

func initBoard(board [3][3]int) [3][3]int {

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = 0
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = getUniqueAndRandomNum(board)
		}
	}
	return board
}

func getNum(board [3][3]int, r int, c int) string {

	if board[r][c] == 9 {
		return " "
	}
	return strconv.Itoa(board[r][c])
}

func isBoardFinished(board [3][3]int) bool {

	if board[0][0] == 1 && board[0][1] == 2 && board[0][2] == 3 &&
		board[1][0] == 4 && board[1][1] == 5 && board[1][2] == 6 &&
		board[2][0] == 7 && board[2][1] == 8 && board[2][2] == 9 {
		return true
	}

	return false
}

func findGap(board [3][3]int) (int, int) { return findPositionOfNum(board, 9) }

func findPositionOfNum(board [3][3]int, num int) (int, int) {

	var positionI, positionJ int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == num {
				positionI = i
				positionJ = j
				goto returning
			}
		}
	}

returning:
	return positionI, positionJ
}

func boardFinished(board [3][3]int, originalBoard [3][3]int) {
	fmt.Println("\n\nThe board is solved.")
	fmt.Println("It took and", totalTimeTaken, "and", stepCounter, "steps to solve")
	fmt.Println("\nThe original board was: ")
	printBoard(originalBoard)
	fmt.Println("\nThe final board is: ")
	printBoard(board)
}

func isBoardSolvable(board [3][3]int) bool {

	var arr [8]int
	var numOfInversions int

	for a, i := 0, 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] != 9 {
				arr[a] = board[i][j]
				a++
			}
		}
	}

	for first := 0; first < len(arr); first++ {
		for second := first + 1; second < len(arr); second++ {
			if arr[second] < arr[first] {
				numOfInversions++
			}
		}
	}

	if numOfInversions%2 == 0 {
		return true
	}

	return false
}
