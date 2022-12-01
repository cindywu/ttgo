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

	board = executeMove(currentMove, board, player)
	fmt.Println("The board is:")
	displayBoard(board)

	// executeMove(moveLocation, player, board)
}

func executeMove(move int, board [9]int, player int) [9]int {
	if move < 10 && board[move] != 0 {
		fmt.Println(move, " is already taken")
		move = makeAMove()
	}

	for move > 9 {
		// TODO: check move is valid int
		fmt.Println("Please enter a valid move")
		move = makeAMove()
	}

	fmt.Println("Move executed")
	return board
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
