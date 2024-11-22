package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Animal struct {
	eats string
	locomotion string
	sound string
}

func (a *Animal) InitAnimal(e,l,s string) {
	a.eats = e
	a.locomotion = l
	a.sound = s
}

func (a *Animal) Eat() {
	fmt.Println(a.eats)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.sound)
}


func main() {
	var cow, bird, snake Animal
	cow.InitAnimal("grass", "walk", "moo")
	bird.InitAnimal("worms", "fly", "peep")
	snake.InitAnimal("mice", "slither", "hsss")

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("Please enter an animal name \"cow\" \"bird\" or \"snake\" and a second word separated by space, which is either \"eat\", \"move\" or \"speak\": ")
		
		if !scanner.Scan() {
				break
		}

		line := scanner.Text()
		text := strings.Fields(line)

		if len(text) != 2 {
			fmt.Println("Please enter words based on the instruction, for example: snake move")
			continue
		}
		
		animals := map[string]*Animal {
			"cow": &cow,
			"snake": &snake,
			"bird": &bird,
		}
		
		
		animal, action := text[0], text[1]

		currentAnimal, exists := animals[animal]

		if !exists {
			fmt.Println("Enter an existing animal")
			continue
		}

		actionName := strings.Title(action)
		method := reflect.ValueOf(currentAnimal).MethodByName(actionName)

		if method.IsValid() {
			method.Call(nil)
		} else {
			snake.Move()
			fmt.Println("Unknown command. Try 'move' or 'eat' for example.")
		}
		fmt.Println("----------------------------------")
	}
}