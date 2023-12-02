package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func mainTwoPartOne() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	redMax := 12
	greenMax := 13
	blueMax := 14
	total := 0
	for scanner.Scan() {
		overLimit := false
		stringParts := strings.Split(scanner.Text(), ":")

		gameIDString := strings.Split(stringParts[0], " ")[1]
		gameID, _ := strconv.Atoi(gameIDString)

		sets := strings.Split(stringParts[1], ";")
		for _, set := range sets {
			setParts := strings.Split(set, ",")
			for i := range setParts {
				if strings.Contains(setParts[i], "red") {
					colourNumString := strings.Replace(setParts[i], "red", "", -1)
					colourNumString = strings.Trim(colourNumString, " ")
					colourNum, _ := strconv.Atoi(colourNumString)
					if colourNum > redMax {
						overLimit = true
					}
				}

				if strings.Contains(setParts[i], "green") {
					colourNumString := strings.Replace(setParts[i], "green", "", -1)
					colourNumString = strings.Trim(colourNumString, " ")
					colourNum, _ := strconv.Atoi(colourNumString)
					if colourNum > greenMax {
						overLimit = true
					}
				}

				if strings.Contains(setParts[i], "blue") {
					colourNumString := strings.Replace(setParts[i], "blue", "", -1)
					colourNumString = strings.Trim(colourNumString, " ")
					colourNum, _ := strconv.Atoi(colourNumString)
					if colourNum > blueMax {
						overLimit = true
					}
				}
			}
		}

		if !overLimit {
			total = total + gameID
		}
	}

	fmt.Printf("Total: %v", total)
}
