package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	lines := splitText()

	re := regexp.MustCompile("[0-9]")

	sum := 0

	for _, line := range lines {
		str := ""

		num := re.FindAllString(line, -1)

		// add first + last numbers in string
		str = str + num[0]
		str = str + num[len(num)-1]

		int, _ := strconv.Atoi(str)

		sum = sum + int
	}

	fmt.Println(sum)

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
