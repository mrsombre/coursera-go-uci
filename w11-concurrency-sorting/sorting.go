package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	threads = 4
)

func readInput(scanner *bufio.Scanner) ([]int, error) {
	fmt.Print("> ")

	scanner.Scan()
	s := scanner.Text()
	if s == "exit" {
		os.Exit(0)
	}

	sa := strings.Split(s, " ")
	var input []int
	for _, strnum := range sa {
		intnum, err := strconv.Atoi(strnum)
		if err != nil {
			return nil, err
		}
		input = append(input, intnum)
	}
	return input, nil
}

func sortInput(input []int) []int {
	fmt.Printf("Input %v\n", input)
	inputSize := len(input)
	// small array
	if inputSize < threads {
		fmt.Printf("Sorting array < %v in 1 thread\n", threads)
		return sortArray(input)
	}
	chunkSize := int(math.Round(float64(inputSize) / threads))
	remain := inputSize - (chunkSize * threads)
	var sorted []int
	mc := make(chan []int, threads)
	for i := 0; i < threads; i++ {
		from := i * chunkSize
		to := (i + 1) * chunkSize
		if to > inputSize {
			to = inputSize
		}
		if remain > 0 && i == threads-1 {
			to += remain
		}
		// do not need sort 1 element
		if to-from == 1 {
			fmt.Printf("Skip sorting 1 element %v\n", input[from:to])
			mc <- input[from:to]
			continue
		}
		go func() {
			mc <- sortArray(input[from:to])
		}()
	}
	for i := 0; i < threads; i++ {
		sorted = sortMerge(sorted, <-mc)
	}
	return sorted
}

func sortArray(a []int) []int {
	fmt.Printf("Sorting %v\n", a)
	sort.Ints(a)
	return a
}

func sortMerge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i := 0
	// smaller first
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	// push remaining numbers
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}

func genExample() string {
	rand.Seed(time.Now().UnixNano())
	var nums []string
	for _, i := range rand.Perm(threads * 2) {
		if rand.Float64() > 0.5 {
			i = -i
		}
		nums = append(nums, strconv.Itoa(i))
	}
	return strings.Join(nums, " ")
}

func main() {
	fmt.Println("Threads sorting:")
	fmt.Println("Enter integers separated by space or type exit")
	fmt.Printf("Example %v\n", genExample())

	scanner := bufio.NewScanner(os.Stdin)
	for {
		input, err := readInput(scanner)
		if err != nil {
			fmt.Printf("%v\n", err.Error())
			continue
		}

		sorted := sortInput(input)
		fmt.Printf("Sorted %v\n", sorted)
	}
}
