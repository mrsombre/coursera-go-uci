package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var data []Name

	var file string
	fmt.Println("Enter a file name:")
	fmt.Scanln(&file)

	_, err := os.Stat(file)
	if err != nil {
		panic(fmt.Sprintf("file %s not exists: %s", file, err.Error()))
	}
	fp, err := os.Open(file)
	if err != nil {
		panic(fmt.Sprintf("error opening a file %s: %s", file, err.Error()))
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var line string
	var lineData []string
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) == 0 {
			continue
		}
		lineData = strings.Split(line, " ")
		if len(lineData) != 2 {
			panic(fmt.Sprintf("wrong line format: %s", line))
		}
		data = append(data, Name{lineData[0], lineData[1]})
	}

	for _, name := range data {
		fmt.Printf("%s %s\n", name.fname, name.lname)
	}
}
