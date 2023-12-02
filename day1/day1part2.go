package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func mainOnePartTwo() {
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
		line = numberifyLine3(line)

		for _, val := range line {
			if contains(validUnicodes, string(val)) {
				digit1 = string(val)
				break
			}
		}

		line = reverseString(line)
		for _, val := range line {
			if contains(validUnicodes, string(val)) {
				digit2 = string(val)
				break
			}
		}

		lineNumber, _ := strconv.Atoi(digit1 + digit2)
		fmt.Printf(" Line Number: %v \n", lineNumber)
		total = total + lineNumber
	}
	fmt.Printf("Total: %v", total)
}

func numberifyLine3(line string) string {
	validNumberStrings := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	validNumbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	numsInLine := ""

	for i := 0; i < len(line); i++ {
		constructedWord := string(line[i])
		_, numErr := strconv.Atoi(constructedWord)
		if numErr == nil {
			numsInLine = numsInLine + constructedWord
			continue
		}

		for c := i + 1; c < len(line); c++ {
			constructedWord = constructedWord + string(line[c])
			for numIndex, numAsString := range validNumberStrings {
				if constructedWord == numAsString {
					numsInLine = numsInLine + validNumbers[numIndex]
				}
			}
		}
	}

	return numsInLine
}
