package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Println("Please, enter a string")
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err.Error())
	}

	input = strings.ToLower(input)

	if input[0:1] == `i` && input[len(input)-1:] == `n` && strings.Contains(input, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
