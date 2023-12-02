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
	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		redLargest := 0
		greenLargest := 0
		blueLargest := 0

		stringParts := strings.Split(scanner.Text(), ":")
		sets := strings.Split(stringParts[1], ";")
		for _, set := range sets {
			setParts := strings.Split(set, ",")
			for i := range setParts {
				if strings.Contains(setParts[i], "red") {
					colourNumString := strings.Replace(setParts[i], "red", "", -1)
					colourNumString = strings.Trim(colourNumString, " ")
					colourNum, _ := strconv.Atoi(colourNumString)
					if colourNum > redLargest {
						redLargest = colourNum
					}
				}

				if strings.Contains(setParts[i], "green") {
					colourNumString := strings.Replace(setParts[i], "green", "", -1)
					colourNumString = strings.Trim(colourNumString, " ")
					colourNum, _ := strconv.Atoi(colourNumString)
					if colourNum > greenLargest {
						greenLargest = colourNum
					}
				}

				if strings.Contains(setParts[i], "blue") {
					colourNumString := strings.Replace(setParts[i], "blue", "", -1)
					colourNumString = strings.Trim(colourNumString, " ")
					colourNum, _ := strconv.Atoi(colourNumString)
					if colourNum > blueLargest {
						blueLargest = colourNum
					}
				}
			}
		}

		total = total + (redLargest * greenLargest * blueLargest)
	}

	fmt.Printf("Total: %v", total)
}
