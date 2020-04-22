package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var input string
	data := map[string]string{}

	fmt.Println("Enter a name:")
	fmt.Scanln(&input)
	data["name"] = input

	fmt.Println("Enter an address:")
	fmt.Scanln(&input)
	data["address"] = input

	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(b))
}
