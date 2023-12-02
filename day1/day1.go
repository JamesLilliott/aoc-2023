package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func mainOnePartOne() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	validUnicodes := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	total := 0
	for scanner.Scan() {
		digit1 := ""
		digit2 := ""

		line := scanner.Text()
		fmt.Println(line)
		for _, val := range line {
			if contains(validUnicodes, string(val)) {
				digit1 = string(val)
				fmt.Println(digit1)
				break
			}
		}

		line = reverseString(line)
		for _, val := range line {
			if contains(validUnicodes, string(val)) {
				digit2 = string(val)
				fmt.Println(digit2)
				break
			}
		}

		fmt.Println(digit1 + digit2)
		lineNumber, _ := strconv.Atoi(digit1 + digit2)
		total = total + lineNumber
		fmt.Println("---")
	}
	fmt.Printf("Total: %v", total)
}
