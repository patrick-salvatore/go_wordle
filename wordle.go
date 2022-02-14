package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

type AttemptedWord struct {
	indexes []int
	word    string
}

var SOLVED = false
var ATTEMPTS = 5
var ATTEMPTED = make([]AttemptedWord, 5)

const SPACE = "  "

const (
	CORRECT = 1 << iota
	MISS
	WRONG
)

func select_word() string {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(Word_Set_1)
	return Word_Set_2[rand.Intn(max-min)+min]
}

func print_empty_line(spacer *color.Color) {
	fmt.Print("\n\n ")
	for i := 0; i < 5; i++ {
		color.New(color.FgBlack).Add(color.BgWhite).Print(SPACE)
		spacer.Print(SPACE)
	}
	fmt.Print("\n\n")
}

func draw(indexes []int, guess string, spacer *color.Color) {
	incorrect_block := color.New(color.FgBlack).Add(color.BgWhite)
	missed_block := color.New(color.FgBlack).Add(color.BgYellow)
	correct_block := color.New(color.FgBlack).Add(color.BgGreen)

	fmt.Print("\n\n ")
	for i, val := range indexes {
		switch val {
		case CORRECT:
			correct_block.Printf(" %v ", string(guess[i]))
		case MISS:
			missed_block.Printf(" %v ", string(guess[i]))
		case WRONG:
			incorrect_block.Printf(" %v ", string(guess[i]))
		}
		spacer.Print(SPACE)
	}
	fmt.Print("\n\n")
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func unquoteCodePoint(s string) (string, error) {
	r, err := strconv.ParseInt(strings.TrimPrefix(s, "\\U"), 16, 32)
	return string(r), err
}

func main() {
	guess_word := select_word()
	var guess string

	spacer := color.New(color.BgBlack)
	fmt.Print("\n Welcome to a golang wordle clone in the terminal \n ")

	for {
		if ATTEMPTS == 0 {
			fmt.Printf("Out of guesses, answer is %v \n", guess_word)
			break
		} else if SOLVED {
			s, _ := unquoteCodePoint("\\U0001f389")
			fmt.Printf(" %s%s Correct, nice job %s%s \n\n", s, s, s, s)
			break
		} else {
			if ATTEMPTS == 5 {
				print_empty_line(spacer)
			}

			fmt.Printf(" %v/5 Input a word: ", 5-ATTEMPTS)
			fmt.Scan(&guess)
			guess = strings.TrimRight(guess, " ")

			if !contains(Word_Set_1, guess) && !contains(Word_Set_2, guess) {
				s, _ := unquoteCodePoint("\\U274c")
				fmt.Printf(" %s%s This is an invalid word %s%s \n\n", s, s, s, s)

			} else {
				new_index := make([]int, len(guess))

				for i, s := range string(guess) {
					gs := rune(guess_word[i])

					if s == gs {
						new_index[i] = CORRECT
					} else if strings.Index(guess_word, string(s)) >= 0 {
						new_index[i] = MISS
					} else {
						new_index[i] = WRONG
					}
				}

				ATTEMPTED[5-ATTEMPTS] = AttemptedWord{indexes: new_index, word: guess}

				c := exec.Command("clear")
				c.Stdout = os.Stdout
				c.Run()

				for _, attempted_word := range ATTEMPTED {
					if len(attempted_word.indexes) > 0 {
						draw(attempted_word.indexes, string(attempted_word.word), spacer)
					}

				}

				ATTEMPTS = ATTEMPTS - 1
			}
		}

		if guess == guess_word {
			SOLVED = true
		}
	}
}
