package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

var randInt int = rand.Intn(100)
var tries int = 5
var guesses int
var name string

func user() {
	fmt.Println("What is your name?")
	fmt.Scanln(&name)
	fmt.Println(name + ", Welcome to the guessing game!")
}

func explain() {
	fmt.Println("I'm thinking of a number between 1 and 100. Try to guess it!")
	fmt.Println("You have " + strconv.Itoa(tries) + " tries to guess the number.")
}

func guess() {
	fmt.Scanln(&guesses)
	if tries == 0 {
		fmt.Println("You ran out of tries. The number was " + strconv.Itoa(randInt))
	} else if guesses == randInt {
		fmt.Println("You guessed the number! It was " + strconv.Itoa(randInt))
	} else if guesses > randInt {
		fmt.Println("Your guess is too high. You have " + strconv.Itoa(tries) + " tries left.")
		tries--
		guess()
	} else if guesses < randInt {
		fmt.Println("Your guess is too low. You have " + strconv.Itoa(tries) + " tries left.")
		tries--
		guess()
	} else if guesses > 100 || guesses < 0 {
		fmt.Println("The number should be between 1 and 100")
	}

}
func restart() {
	var choice string
	fmt.Println("Would you like to play again?(y/n)")
	fmt.Scanln(&choice)
	if choice == "y" {
		fmt.Println("Restarting...")
		randInt = rand.Intn(100)
		tries = 5
		guess()
	} else if choice == "n" {
		fmt.Println("Thanks for playing!")
	} else {
		fmt.Println("Invalid choice")
		restart()
	}
}

func main() {
	user()
	explain()
	guess()
	restart()
}
