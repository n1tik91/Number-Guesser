# Number Guessing Game 🎲

A simple command-line game written in Go where the user guesses a randomly generated number between 1 and 100. This project is designed for beginners learning Go programming.



## 🚀 Features

- Generates a random number between 1 and 100.
- Prompts the user to guess the number.
- Provides feedback ("Too high!" or "Too low!") for each guess.
- Tracks the number of attempts.
- Congratulates the user upon a correct guess.

---

## 🛠️ Requirements

- **Go**: [Install Go](https://golang.org/dl/) (any recent version will work).

---

## 💻 How to Run

1. **Clone the Repository**:
    ```bash
    git clone https://github.com/n1tik91/Number-Guesser.git
    cd Number-Guesser
    ```

2. **Run the Program**:
    ```bash
    go run main.go
    ```

3. **Play the Game**:
    - Enter your guesses when prompted.
    - Follow the feedback until you guess the number.

---

## 🧩 Example Output

```plaintext
Welcome to the Number Guessing Game!
I am thinking of a number between 1 and 100.
Enter your guess: 50
Too high! Try again.
Enter your guess: 25
Too low! Try again.
Enter your guess: 37
Congratulations! You guessed it in 3 attempts.
Thank you for playing!
