package main

// import "fmt"
// import "math"
// import "math/rand"

// Better yet, you can import multiple packages in a single line:
import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("You want to play?")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

// func displayBoard(board [9]) {

// }
