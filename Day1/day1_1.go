package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("1_1_input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var increase int = 0
	var check_first_line bool = true
	var temp_below int

	for scanner.Scan() {

		temp_above, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if check_first_line == true {
			check_first_line = false
		} else {
			if temp_above > temp_below {
				increase = increase + 1
			}
		}
		temp_below = temp_above
	}
	fmt.Println(increase)
}
