package main

import "fmt"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File [8]bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	countFile := 0
	for _, v := range cb[file] {
		if v {
			countFile += 1
		}
	}

	return countFile
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	countRank := 0

	if rank >= 1 && rank <= 8 {
		for k, _ := range cb {
			if cb[k][rank-1] {
				countRank += 1
			}
		}

		return countRank
	}

	return countRank
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	countBoard := 0

	for _, v := range cb {
		countBoard += len(v)
	}

	return countBoard
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	countOccupy := 0

	for k, _ := range cb {
		for _, v := range cb[k] {
			if v {
				countOccupy += 1
			}
		}
	}

	return countOccupy
}

func main() {
	var board = Chessboard{
		"A": {true, false, true, false, true, false, false, false},
		"B": {false, false, false, false, false, false, false, false},
		"C": {true, false, false, false, false, false, false, false},
		"D": {true, true, false, false, false, false, false, false},
		"E": {false, false, false, false, false, false, false, false},
		"F": {false, false, true, true, false, false, false, false},
		"G": {false, false, false, false, false, false, false, false},
		"H": {false, false, false, false, false, true, false, false},
	}

	fmt.Println(CountInFile(board, "A"))
	// => 3

	fmt.Println(CountInRank(board, 2))
	// => 1

	fmt.Println(CountAll(board))
	// => 64

	fmt.Println(CountOccupied(board))
	// => 15
}
