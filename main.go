package main

import (
	"encoding/csv"
	"strings"
	"bufio"
	"time"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Open and read the csv file
	fileName := flag.String("csv", "./questions.csv", "path to questions csv")

	flag.Parse()

	questionsFile, err := os.Open(*fileName)

	if err != nil {
		fmt.Println("An error encountered ::", err)
		os.Exit(1)
	}

	csvReader := csv.NewReader(questionsFile)
	questions, err := csvReader.ReadAll()
	// 

	// Initiialize the console input reader
	consoleReader := bufio.NewReader(os.Stdin)
	// 

	var score int
	fmt.Println("Please answer the following questions:")

	// Initiate Timer
	const seconds = 20

	time.AfterFunc(time.Second * seconds, func() {
		fmt.Println("Times up!")
	  fmt.Printf("You got %d out of %d questions correct.", score, len(questions))
		os.Exit(1)
	})
	// 

	// Iterate through each question asking questions and updating score
	for i := range questions {
		question := questions[i][0]
		answer := questions[i][1]

		fmt.Printf("Question #%d: %s\n", i + 1, question)

		text, _ := consoleReader.ReadString('\n')
		userAnswer := strings.TrimSpace(text)

		if userAnswer == answer {
			score += 1
		}
	}
	//

	fmt.Printf("You got %d out of %d questions correct.", score, len(questions))
	os.Exit(1)
}