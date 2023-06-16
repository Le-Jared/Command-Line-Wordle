package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var wordList = []string{
	"apple", "table", "grape", "chair", "place", "bread", "light", "sound", "earth",
}

type GuessResult struct {
	CorrectPosition int
	CorrectLetter   int
}

func evaluateGuess(target, guess string) GuessResult {
	result := GuessResult{}

	for i := 0; i < len(target); i++ {
		if target[i] == guess[i] {
			result.CorrectPosition++
		}
	}

	for _, letter := range guess {
		if strings.Contains(target, string(letter)) && !strings.Contains(guess[:strings.Index(guess, string(letter))], string(letter)) {
			result.CorrectLetter++
		}
	}

	result.CorrectLetter -= result.CorrectPosition
	return result
}

func isWordInList(word string, list []string) bool {
	for _, b := range list {
		if b == word {
			return true
		}
	}
	return false
}

func main() {
    s := rand.NewSource(time.Now().Unix())
    r := rand.New(s)
    target := wordList[r.Intn(len(wordList))]

	scanner := bufio.NewScanner(os.Stdin)

	guesses := make([]string, 0)

	for i := 0; i < 6; i++ {
		fmt.Printf("Guess #%d: ", i+1)
        scanner.Scan()
        guess := strings.ToLower(scanner.Text())

		if len(guess) != 5 {
			fmt.Println("Please input a 5-letter word.")
			i-- // decrement i to retry the same turn
			continue
		}

		if !isWordInList(guess, wordList) {
			fmt.Println("Your guess is not a valid word. Please try again.")
			i-- // decrement i to retry the same turn
			continue
		}

		if isWordInList(guess, guesses) {
			fmt.Println("You've already guessed this word. Please try again.")
			i-- // decrement i to retry the same turn
			continue
		}

		guesses = append(guesses, guess)

		result := evaluateGuess(target, guess)

		if result.CorrectPosition == 5 {
			fmt.Println("Congrats! You've guessed the correct word!")
			return
		}

		if i == 2 {
			fmt.Printf("Hint! The letter at position 1 is %c.\n", target[0])
		}

		fmt.Printf("Correct letters: %d, Letters in right position: %d\n", result.CorrectLetter, result.CorrectPosition)
	}

	fmt.Printf("You didn't guess the word. It was %s.\n", target)
}
