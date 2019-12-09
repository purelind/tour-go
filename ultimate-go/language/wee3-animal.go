//
// week3 module3 animal.go
//

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Animal struct {
	food		string
	locomotion	string
	noise		string
}

func main() {
	cow := Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
	bird := Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}
	snake := Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}

	for {
		fmt.Println("\n[ cow | bird | snake ]  [ eat | move | speak ]")
		fmt.Print("Please enter animal and requestInfo, split them with white space: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		line := scanner.Text()
		words := strings.Fields(line)
		ok := verifyInput(words)
		if !ok {
			log.Println("Invalid input")
			continue
		}

		switch words[0] {
		case "cow":
			if words[1] == "eat" {
				cow.Eat()
			} else if words[1] == "move" {
				cow.Move()
			} else {
				cow.Speak()
			}
		case "bird":
			if words[1] == "eat" {
				bird.Eat()
			} else if words[1] == "move" {
				bird.Move()
			} else {
				bird.Speak()
			}
		case "snake":
			if words[1] == "eat" {
				snake.Eat()
			} else if words[1] == "move" {
				snake.Move()
			} else {
				snake.Speak()
			}
		}
	}
}

func (a Animal) Eat() {
	fmt.Println(" -> " + a.food)
}

func (a Animal)  Move() {
	fmt.Println(" -> " + a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(" -> " + a.noise)
}

func verifyInput(a []string) bool {
	validAnimal := false
	validMethod := false
	for _, name := range []string{ "cow", "bird", "snake" } {
		if a[0] == name {
			validAnimal = true
		}
	}
	for _, name := range []string{ "eat", "move", "speak" } {
		if a[1] == name {
			validMethod = true
		}
	}

	return validAnimal && validMethod
}