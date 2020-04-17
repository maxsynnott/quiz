package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"bufio"
	"strings"
)

func main() {
	// Open and read the csv file
	questionsFile, err := os.Open("./questions.csv")
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}

	csvReader := csv.NewReader(questionsFile)
	questions, err := csvReader.ReadAll()
	// 

	// Initiialize the console input reader
	consoleReader := bufio.NewReader(os.Stdin)
	// 

	// Iterate through each question
	fmt.Println("Please answer the following questions:")

	var score int

	for i := range questions {
		question := questions[i][0]
		answer := questions[i][1]

		fmt.Printf("Question #%d: %s\n", i, question)

		text, _ := consoleReader.ReadString('\n')
		userAnswer := strings.TrimSpace(text)

		if userAnswer == answer {
			score += 1
		}
	}
	//

	fmt.Printf("You got %d out of %d questions correct.", score, len(questions))
}