## Command Line Wordle
Command-Line Wordle is a text-based adaptation of the popular online game Wordle, designed to be played in your terminal or command line interface. The game will generate a random 5-letter word from a predefined list, and your goal is to guess that word in six attempts or fewer.

Each time you guess, the game will tell you how many letters in your guess are in the target word (but in the wrong position), and how many letters are in the correct position.

## How to Play
To start the game, run the wordle.go file in your Go environment.

When prompted, type in a 5-letter word from the terminal and hit Enter to submit your guess. If your guess is not valid (e.g. it's not a 5-letter word, or it's not in the word list), you'll be asked to guess again.

After each guess, you'll be told how many letters in your guess are in the correct position in the target word, and how many additional letters are in the target word but in the wrong position.

If you guess the target word within six attempts, you win! If not, the game will tell you what the target word was.
