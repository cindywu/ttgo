package main

import (
	"fmt"
)

func main() {
	board := [9]string{"", "", "", "", "", "", "", "", ""}
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

func executeMove(move int, board [9]string, player int) [9]string {
	if move < 10 && board[move] != "" {
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

func displayBoard(board [9]string) {
	for i, v := range board {
		if v == "" {
			fmt.Printf("%d", i+1)
		}

		if (i+1)%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf(" | ")
		}
	}
}
