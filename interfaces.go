package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"errors"
)



type Animal interface {
	Move()
	Eat()
	Speak()
}

type Cow struct{
	food string
	locomotion string
	sound string
	name string
}

func (a *Cow) Move() {
	fmt.Println(a.locomotion)
}

func (a *Cow) Eat() {
	fmt.Println(a.food)
}

func (a *Cow) Speak() {
	fmt.Println(a.sound)
}

type Snake struct{
	food string
	locomotion string
	sound string
	name string
}

func (a *Snake) Move() {
	fmt.Println(a.locomotion)
}

func (a *Snake) Eat() {
	fmt.Println(a.food)
}

func (a *Snake) Speak() {
	fmt.Println(a.sound)
}

type Bird struct{
	food string
	locomotion string
	sound string
	name string
}

func (a *Bird) Move() {
	fmt.Println(a.locomotion)
}

func (a *Bird) Eat() {
	fmt.Println(a.food)
}

func (a *Bird) Speak() {
	fmt.Println(a.sound)
}


func NewAnimal(name string, animalType string) (Animal, error){
	switch animalType {
		case "cow":
			return &Cow{
				food:       "grass",
				locomotion: "walk",
				sound:      "moo",
				name: name,
			}, nil
		case "bird":
			return &Bird{
				food:       "worms",
				locomotion: "fly",
				sound:      "peep",
				name: name,
			}, nil
		case "snake":
		return &Snake{
			food:       "mice",
			locomotion: "slither",
			sound:      "hsss",
			name: name,
		}, nil
		default:
			return nil, errors.New("Unknown animal: use snake, bird or cow")
	}
}

func main() {
	animals := make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter: newanimal name type OR query name action (exit for ending the program)>")
		scanner.Scan()
		line := scanner.Text()
		text := strings.Fields(line)
		if line == "exit" {
			fmt.Println("Exiting program...")
			break
		}
		if len(text) != 3 {
			fmt.Println("You must enter three words")
			continue
		}

		input1, input2, input3 := text[0], text[1], text[2]

		if input1 == "newanimal" {
			if _, exists := animals[input2]; exists {
				fmt.Println("Animal already exists")
				continue
			}

			animal, inputError := NewAnimal(input2, input3)
			if inputError != nil {
				fmt.Println(inputError)
				continue
			}
			animals[input2] = animal
			fmt.Println("Animal succesfully added to list")

		} else if input1 == "query" {
			animal, exists := animals[input2]
			if !exists {
				fmt.Println("Animal with this name doesn't exist")
				continue
			}
			switch input3 {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "sound":
				animal.Speak()
			default:
				fmt.Println(errors.New("Invalid action for animal"))
				continue
			}
		} else {
			fmt.Println(errors.New("Unknown command: use newanimal or query"))
		}
	}
}