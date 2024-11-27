package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	// Generate a random number between 1 and 100
	targetNumber := rand.Intn(100) + 1

	var userGuess int
	var attempts int

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I have picked a number between 1 and 100.")
	fmt.Println("Can you guess what it is?")

	for {
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scan(&userGuess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 100.")
			continue
		}

		attempts++

		// Check the user's guess
		if userGuess < targetNumber {
			fmt.Println("Too low! Try again.")
		} else if userGuess > targetNumber {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congratulations! You guessed the number in %d attempts.\n", attempts)
			break
		}
	}

	fmt.Println("Thanks for playing!")
}
