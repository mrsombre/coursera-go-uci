package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := make([]int, 0, 3)
	var dataResult []string

	var input string
	for {
		fmt.Println("Type an integer or X to finish:")

		fmt.Scanln(&input)
		num, err := strconv.Atoi(input)
		if err != nil {
			if strings.EqualFold(input, "x") {
				return
			}
			fmt.Println("Only integers or X allowed!")
			continue
		}

		data = append(data, num)
		sort.Ints(data)

		dataResult = make([]string, len(data))
		for i, v := range data {
			dataResult[i] = strconv.Itoa(v)
		}
		fmt.Println(strings.Join(dataResult, ", "))

		input = ""
	}
}
