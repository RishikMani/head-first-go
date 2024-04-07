package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func randInt(min, max int) int {
  // rand.Intn generates number from [0, n) only. We need it to generate from [1,n].
	return min + rand.Intn(max-min)
}

func main() {
	randomNumber := randInt(1, 100)
	for i := 1; i < 11; i++ {
		fmt.Print("Enter the number you just guessed: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal((err))
		}
		input = strings.TrimSpace(input)
		intInput, err := strconv.Atoi(input)
		if intInput < randomNumber {
			fmt.Println("OOPS. Your guess was LOW.")
		} else if intInput == randomNumber {
			fmt.Println("Good job! You guessed it.")
			break
		} else {
			fmt.Println("OOPS. Your guess was HIGH.")
		}
		fmt.Println("You have", 10-i, "chances left to guess the number.")

		if i == 10 && intInput != randomNumber {
			fmt.Println("Sorry. You didn't guess my number. It was:", randomNumber)
		}
	}
}
