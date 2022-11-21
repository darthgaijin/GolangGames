package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var plyLives = 8

// Generate random number between 1 and the length of the current dictionary of words.
func selectRandWord() string {
	var dictionary = []string{
		"Rick Sanchez",
		"Hello World",
		"This Game Sucks",
		"Invention",
	}

	var dicLength int = len(dictionary)
	randNum := rand.Intn(dicLength)
	return dictionary[randNum]
}

// Display the word that is in question but this is blanked out by _ so that we know what has been guessed and what has not
// This logic takes in the word that is generated and spits out a ____ ___ ___ version of the word.
func displayWord(gameWord string) string {
	chars := []rune(gameWord)
	var blankWords string
	for i := 0; i < len(chars); i++ {
		stringChar := string(chars[i])
		if stringChar != " " {
			blankWords += "_"
		} else {
			blankWords += " "
		}
	}

	return blankWords
}

// Logic to see if the guess was correct
func playersGuess(word string, playerInput string) bool {
	result := strings.Contains(word, playerInput)
	return result
}

// Take in the correct guess by the user and do something with it.
// This will take in the guess compaire it to the word and then act accordingly.
// Player guess will then populate the blanked word with the correctly guessed letters.
func playerGuessCalculated(word *string, playerInput *string, currentBlank string) string {
	chars := []rune(*word)
	blankChars := []rune(currentBlank)
	var blankWords string
	for i := 0; i < len(chars); i++ {
		stringChar := string(chars[i])
		blankstringChar := string(blankChars[i])
		if stringChar == *playerInput {
			blankWords += *playerInput
		} else {
			blankWords += blankstringChar
		}
	}

	return blankWords
}

func main() {
	word := selectRandWord()
	fmt.Println("Welcome to this janky game!")
	fmt.Println("===========================")
	fmt.Println("Your word is below:")
	blankWords := displayWord(word)
	fmt.Println(word)
	fmt.Println(blankWords)

	for {
		var playerInput string

		if plyLives == 1 {
			fmt.Println("Last life left. Make a word guess!")
			fmt.Scanln(&playerInput)
			fmt.Println("Your typed:", playerInput)
			if word != playerInput {
				fmt.Println("Wrong! Game Over!")
				break
			}
			fmt.Println("Wrong! Game Over!")

		} else {
			fmt.Println("Lives Left: ", plyLives)
			fmt.Print("Type a Letter: ")
			fmt.Scanln(&playerInput)
			fmt.Println("Your typed:", playerInput)
			if playersGuess(word, playerInput) {
				fmt.Println("Success its on there.")
				blankWords = playerGuessCalculated(&word, &playerInput, blankWords)
				fmt.Println(blankWords)
			} else {
				plyLives -= 1
				if plyLives == 0 {
					break
				}
			}
			if !strings.Contains(blankWords, "_") {
				fmt.Println("Game Won.")
				break
			}
		}
	}
}
