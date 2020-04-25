package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Type an integers sequence separated with spaces:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	inputInts := strings.Split(input, " ")
	fmt.Printf("Your input: %v\n", inputInts)

	slice := make([]int, len(inputInts))
	for i, v := range inputInts {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("one of input values is not an integer: %v", v)
			os.Exit(1)
		}
		slice[i] = num
	}

	BubbleSort(slice)
	assertSorted(slice)

	fmt.Printf("Sorted slice: %v", slice)
}

func BubbleSort(slice []int) {
	sorted := true
	cnt := len(slice) - 1
	for i := 0; i < cnt; i++ {
		sorted = true
		for j := 0; j < cnt-i; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
				sorted = false
			}
		}
		if sorted {
			fmt.Printf("No need further sorting on iteration %d of %d\n", i, cnt)
			break
		}
	}
}

func Swap(slice []int, i int) {
	to := slice[i+1]
	slice[i+1] = slice[i]
	slice[i] = to
}

func assertSorted(slice []int) {
	if !sort.IntsAreSorted(slice) {
		fmt.Printf("slice is not sorted\n")
		os.Exit(1)
	}
}
