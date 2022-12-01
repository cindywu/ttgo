package main

import (
	"fmt"
)

func main() {
	board := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	player := 1

	fmt.Println("You want to play?")
	fmt.Println("The board is:")
	displayBoard(board)
	fmt.Println("Player ", player, " plays:")
	currentMove := makeAMove()

	executeMove(currentMove)

	// executeMove(moveLocation, player, board)
}

func executeMove(move int) {
	fmt.Println("Move executed")
}

func makeAMove() int {
	var moveLocation int
	fmt.Scan(&moveLocation)
	return moveLocation
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
