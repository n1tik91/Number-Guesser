
# Number Guessing Game 🎲

A fun command-line game written in Go where the user tries to guess a randomly generated number. The program provides feedback for each guess and tracks the number of attempts.

---

## 📖 Overview
The **Number Guessing Game** generates a random number between 1 and 100. Players are prompted to guess the number, and the program responds with hints like "Too high!" or "Too low!" until the correct number is guessed. The game tracks how many attempts it took the player to guess correctly.

---

## 🚀 Features
- **Random Number Generation**: A new random number is generated for every game session.
- **User Input**: Players enter their guesses via the command line.
- **Feedback Mechanism**: Provides hints to guide the player (e.g., "Too high!" or "Too low!").
- **Attempt Tracker**: Displays the total number of attempts when the game ends.
- **Replayable**: Restart the program to play again.

---

## 🛠️ Requirements
- **Go**: [Install Go](https://golang.org/dl/) (v1.16 or later is recommended).

---

## 💻 How to Run
1. **Clone the Repository**:
    ```bash
    git clone https://github.com/your-username/number-guessing-game.git
    cd number-guessing-game
    ```

2. **Run the Program**:
    ```bash
    go run main.go
    ```

3. **Gameplay**:
    - The program will prompt you to guess a number.
    - Enter your guesses and follow the feedback until you guess correctly.
    - Enjoy!

---

## 🧩 Code Explanation
Here’s an overview of how the program works:
1. **Random Number Generation**: Uses `math/rand` with a seed based on the current time to ensure randomness.
2. **User Input**: Reads player guesses using `fmt.Scan`.
3. **Feedback Loop**: Continuously provides feedback ("Too high!" or "Too low!") until the correct guess is made.
4. **Exit and Summary**: Displays the total number of attempts once the game ends.

---

## 🛠️ Example Output
```plaintext
Welcome to the Number Guessing Game!
I have picked a number between 1 and 100.
Can you guess what it is?

Enter your guess: 50
Too high! Try again.

Enter your guess: 25
Too low! Try again.

Enter your guess: 37
Congratulations! You guessed the number in 3 attempts.
Thanks for playing!
