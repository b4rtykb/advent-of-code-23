package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines := splitText()

	for k, v := range lines {
		fmt.Println(k, v)
	}
}

func splitText() []string {
	file, err := os.Open("./inputs/day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}
