package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	randomNumber := rand.Intn(100) // 0~99
	maxGuesses := 7
	cnt := 0
	for {
		if cnt >= maxGuesses {
			fmt.Printf("Game over! The number was %d.\n", randomNumber)
			break
		}

		fmt.Printf("Guess %d, Please enter number:\n", cnt+1)
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("That's not a valid number. Please try again.")
			continue
		} else {
			if n == randomNumber {
				fmt.Printf("Correct! You got it in %d guesses.\n", cnt+1)
				break
			} else if n > randomNumber {
				fmt.Println("Lower")
			} else {
				fmt.Println("Upper")
			}
			cnt++
		}
	}

}
