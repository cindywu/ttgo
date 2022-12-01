package main

import (
	"fmt"
)

func main() {
	var board = [9]bool{false, false, false, false, false, false, false, false, false}
	board2 := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println("You want to play?")
	fmt.Println("The board is:")
	fmt.Println(board)
	fmt.Println("The board2 is:")
	fmt.Println(board2)
}

// func displayBoard(board [9]) {

// }
