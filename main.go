package main

import (
	"fmt"
)

func main() {
	board := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	fmt.Println("You want to play?")
	fmt.Println("The board is:")
	displayBoard(board)
}

func displayBoard(board [9]int) {
	for i, v := range board {
		if v == 0 {
			fmt.Printf("%d", i)
		}

		if (i+1)%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf(" | ")
		}
	}
}
