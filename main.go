package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

/*
  first we need to read from the words.txt file and select a word randomly
*/
func NewWords() []string {
	data, err := os.ReadFile("words.txt")

	str := string(data)
	str = strings.TrimSpace(str)
	if err != nil {
		log.Fatal("Could not read words.txt file")
	}
	words := strings.Split(str, "\n")
	return words
}

func SelectWord(words []string) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	i := r.Intn(len(words))
	return words[i]
}

func HasWord(word string, words []string) bool {
	var iterations = 0
	i, j := 0, len(words)
	for i + 1< j {
		if iterations >= 100 {
			log.Fatal("infinite loop")
		}
		iterations++
		m := (i + j) / 2
		if words[m] == word {
			return true
		} else if words[m] < word {
			i = m
		} else {
			j = m
		}
	}
	return false
	
}
/*
  wordle a fun word game
	1. users must enter words found on on the words.txt file
	2. if the user has found a letter in the unknown word, and the position is correct, the letter should be green
	3. if the user has found a letter in the unknown word, but the position is incorrect, the letter is yelllow
	4. if the user has failed 6 times then the game is over and the word revealed
	5. if the user successfully guesses the word then the game is over and the word revealed
	6. if the letter is not in the word, then the color is gray
*/
func main() {
	// Clean() <-- run this if receiving a new words_alpha.txt file
	fmt.Println("Welcome to wordle")
	words := NewWords()
	word := SelectWord(words)
	letters := make(map[byte]bool)
	yellow := color.New(color.FgYellow).PrintFunc()
	blue := color.New(color.FgBlue).PrintFunc()
	green := color.New(color.FgGreen).PrintFunc()
	red := color.New(color.FgRed).PrintFunc()
	for i := range word {
		b := word[i]
		letters[b] = true
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("guess a word")
	fmt.Println("-----")
	var won bool
	for i := 0; i < 6; i++ {
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("invalid input")
			log.Fatal(err)
		}
		text = strings.TrimSpace(text)
		if len(text) != len(word) {
			fmt.Printf("submit a %d letter word\n", len(word))
			i--
			continue
		}
		if HasWord(text, words) == false {
			fmt.Printf("%s is not in word in the scrabble dictionary, submit another\n", text)
			i--
			continue
		}

		var correct int
		for i := range word {
			correctLetter, userLetter := word[i], text[i]
			if correctLetter == userLetter {
				green(string(userLetter))
				correct++
			} else if letters[userLetter] {
				yellow(string(userLetter))
			} else {
				blue(string(userLetter))
			}
		}
		fmt.Print("\n")
		if correct == len(word) {
			won = true
			break
		} 
	}
	if won == false {
		red(word)
		fmt.Print("\n")
		fmt.Println("Better luck next time!")
	} else {
		fmt.Println("Nice! You guessed it!")
	}
}