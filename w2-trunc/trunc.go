package main

import (
	"fmt"
)

func main() {
	var userNum float64

	fmt.Println("Please, enter a float number (like 12.345)")
	_, err := fmt.Scanln(&userNum)
	if err != nil {
		panic(err.Error())
	}

	var intNum int
	intNum = int(userNum)

	fmt.Printf("Truncated number: %d, original number: %g", intNum, userNum)
}
