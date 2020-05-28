package main

import (
	"bufio"
	"fmt"
	"os"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (a *Cow) Eat() {
	fmt.Println("grass")
}
func (a *Cow) Move() {
	fmt.Println("walk")
}
func (a *Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (a *Bird) Eat() {
	fmt.Println("worms")
}
func (a *Bird) Move() {
	fmt.Println("fly")
}
func (a *Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (a *Snake) Eat() {
	fmt.Println("mice")
}
func (a *Snake) Move() {
	fmt.Println("slither")
}
func (a *Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	fmt.Println("Type an animal name and an query space separated or word exit")

	var animalName, actionName string
	var animal Animal
	b := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for {
		fmt.Print("> ")
		b.Scan()
		if b.Text() == "exit" {
			os.Exit(0)
		}

		num, err := fmt.Sscan(b.Text(), &animalName, &actionName)
		if num != 2 {
			fmt.Println("Expected two words in the request")
			continue
		}
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		switch animalName {
		case "cow":
			animal = &Cow{}
			break
		case "bird":
			animal = &Bird{}
			break
		case "snake":
			animal = &Snake{}
			break
		default:
			fmt.Printf("Not found an animal with a name: %s\n", animalName)
			continue
		}

		switch actionName {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Printf("No information how %s %s\n", animalName, actionName)
			continue
		}
	}
}
