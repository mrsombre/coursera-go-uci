package main

import (
	"bufio"
	"fmt"
	"os"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() string {
	return a.food
}

func (a *Animal) Move() string {
	return a.locomotion
}

func (a *Animal) Speak() string {
	return a.noise
}

func NewAnimal(food, locomotion, noise string) *Animal {
	return &Animal{
		food:       food,
		locomotion: locomotion,
		noise:      noise,
	}
}

func main() {
	animals := map[string]*Animal{
		"cow":   NewAnimal("grass", "walk", "moo"),
		"bird":  NewAnimal("worms", "fly", "peep"),
		"snake": NewAnimal("mice", "slither", "hsss"),
	}

	var result string
	var animalName, actionName string
	b := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for {
		fmt.Print("> ")
		b.Scan()
		num, err := fmt.Sscan(b.Text(), &animalName, &actionName)
		if num != 2 {
			fmt.Println("Expected two words in the request")
			os.Exit(1)
		}
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		animal, found := animals[animalName]
		if !found {
			fmt.Printf("Not found a with name: %s", animalName)
			os.Exit(1)
		}

		switch actionName {
		case "eat":
			result = animal.Eat()
		case "move":
			result = animal.Move()
		case "speak":
			result = animal.Speak()
		default:
			fmt.Printf("Not found information about: %s", actionName)
			os.Exit(1)
		}

		fmt.Println(result)
	}
}
