package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2_1_input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var x, aim, depth int = 0, 0, 0

	for scanner.Scan() {

		split := strings.Split(scanner.Text(), " ")
		command := split[0]
		value, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
			return
		}

		if command == "up" {
			aim = aim - value
		} else if command == "down" {
			aim = aim + value
		} else {
			x = x + value
			depth = depth + aim*value
		}

	}
	fmt.Println(x * depth)
}
