// a simple game where you need to guess a random number from 1 to 100
package guess

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

//starts the game
func Guess() {
	var numberToGuess int = rand.Intn(101)
	totalGuessesAmount := 10
	victoryStatus := false
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello! Let's play a little game.")
	fmt.Println("I'll guess one number in range of 0-100 and you will try to guess it!")
	fmt.Println("You will have only", totalGuessesAmount, "tries.")
	fmt.Println("Let's start!")
	for guessCounter := 0; guessCounter < totalGuessesAmount; guessCounter++ {
		fmt.Println("\nYou have got", totalGuessesAmount-guessCounter, "chances to guess left.")
		fmt.Print("Your guess number ", guessCounter+1, " is: ")
		inputStr, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		parsedIntNum, err := strconv.Atoi(strings.TrimSpace(inputStr))
		if err != nil {
			log.Fatal(err)
		}
		if parsedIntNum == numberToGuess {
			fmt.Println("YOU WON! ggwp")
			victoryStatus = true
			break
		} else if parsedIntNum > numberToGuess {
			fmt.Println("Oops. Your guess was too HIGH.")
		} else {
			fmt.Println("Oops. Your guess was too LOW.")
		}
	}
	if !victoryStatus {
		fmt.Println("Sorry, you didn't guess my number in", totalGuessesAmount, "tries. It was:", numberToGuess)
	}
}

