package main

import (
	"fmt"
)

func main() {
	board := [9]string{"", "", "", "", "", "", "", "", ""}
	// winCombos := [8][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// 	{1, 4, 7},
	// 	{2, 5, 8},
	// 	{3, 6, 9},
	// 	{1, 5, 9},
	// 	{3, 5, 7}}

	player := "X"

	// win := false

	fmt.Println(`tic tac toe, three in a row...`)
	displayBoard(board)
	fmt.Println(player, "'s turn! pick a spot...")
	currentMove := makeAMove()
	board = executeMove(currentMove, board, player)
	displayBoard(board)

}

func switchPlayer(player string) string {
	if player == "O" {
		player = "X"
	} else {
		player = "O"
	}
	return player
}

func checkIfWin(win bool, board [9]string, winCombos [8][3]int) bool {
	for i := 0; i < len(winCombos); i++ {
		win = false
	}
	return win
}

func executeMove(move int, board [9]string, player string) [9]string {
	if move < 10 && board[move] == "" {
		fmt.Println(player, "picked", move)
		board[move-1] = player
	} else if board[move] != "" {
		fmt.Println("That spot is taken, pick another!")
	}

	for move > 9 {
		fmt.Println("Please enter a valid move")
		move = makeAMove()
	}

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
		} else {
			fmt.Printf(v)
		}

		if (i+1)%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf(" | ")
		}
	}
}
