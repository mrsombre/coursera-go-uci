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

type Cow struct {
	name string
}

func (a *Cow) Eat() {
	fmt.Println("grass")
}
func (a *Cow) Move() {
	fmt.Println("walk")
}
func (a *Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
	name string
}

func (a *Bird) Eat() {
	fmt.Println("worms")
}
func (a *Bird) Move() {
	fmt.Println("fly")
}
func (a *Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
	name string
}

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
	fmt.Println("Create a new animal: newanimal <Name> <Type>")
	fmt.Println("Query an animal info: query <Name> <Action>")

	var command, param1, param2 string
	var animal Animal
	animals := map[string]Animal{}
	b := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for {
		fmt.Print("> ")
		b.Scan()
		// exit
		if b.Text() == "exit" {
			os.Exit(0)
		}

		num, err := fmt.Sscan(b.Text(), &command, &param1, &param2)
		if num != 3 {
			fmt.Println("Expected three words in the request")
			continue
		}
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		switch command {
		case "newanimal":
			if _, found := animals[param1]; found {
				fmt.Printf("Animal with name %s already exists\n", param1)
				continue
			}
			switch param2 {
			case "cow":
				animal = &Cow{param1}
				break
			case "bird":
				animal = &Bird{param1}
				break
			case "snake":
				animal = &Snake{param1}
				break
			default:
				fmt.Printf("Not found an animal type %s\n", param2)
				continue
			}
			animals[param1] = animal
			break
		case "query":
			if _, found := animals[param1]; !found {
				fmt.Printf("Not found %s animal\n", param1)
				continue
			}
			animal = animals[param1]
			switch param2 {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Printf("Not found animal action %s\n", param2)
				continue
			}
			break
		default:
			fmt.Printf("Command %s not supported\n", command)
			continue
		}
	}
}
