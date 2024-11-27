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
	secretNumber := rand.Intn(100) + 1

	// Welcome message
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I am thinking of a number between 1 and 100.")

	// Variables for user guess and attempt count
	var guess int
	attempts := 0

	// Game loop
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess) // Read the user's guess
		attempts++       // Count the attempt

		// Check if the guess is correct
		if guess < secretNumber {
			fmt.Println("Too low! Try again.")
		} else if guess > secretNumber {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congratulations! You guessed it in %d attempts.\n", attempts)
			break // Exit the loop when guessed correctly
		}
	}

	// Goodbye message
	fmt.Println("Thank you for playing!")
}
