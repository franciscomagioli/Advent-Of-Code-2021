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
	var check_third_line int = 1
	var temp_minus_three, temp_minus_two, temp_minus_one int //could use an array size = 4

	for scanner.Scan() {

		temp_now, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if check_third_line == 3 {
			check_third_line = check_third_line + 1
		} else {
			if temp_minus_three != 0 { //sums and validate
				sum_now := temp_now + temp_minus_one + temp_minus_two
				sum_old := temp_minus_one + temp_minus_two + temp_minus_three
				if sum_now > sum_old {
					increase = increase + 1
				}
			}
		}
		temp_minus_three = temp_minus_two
		temp_minus_two = temp_minus_one
		temp_minus_one = temp_now
	}
	fmt.Println(increase)
}
