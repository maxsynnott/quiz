package main

import (
	"encoding/csv"
	"math/rand"
	"strings"
	"bufio"
	"time"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Parse flags
	fileName := flag.String("csv", "./questions.csv", "path to questions csv")
	seconds := flag.Int("seconds", 30, "timer length in seconds")
	shuffle := flag.Bool("shuffle", false, "should questions be shuffled?")

	flag.Parse()
	// 

	// Open and read the csv file
	questionsFile, err := os.Open(*fileName)

	if err != nil {
		fmt.Println("An error encountered ::", err)
		os.Exit(1)
	}

	csvReader := csv.NewReader(questionsFile)
	questions, err := csvReader.ReadAll()
	// 

	// Shuffle the questions if flag provided
	if *shuffle {
		rand.Seed(time.Now().UnixNano())

		rand.Shuffle(len(questions), func(i, j int) {
			questions[i], questions[j] = questions[j], questions[i]
		})
	}
	// 

	// Initiialize the console input reader
	consoleReader := bufio.NewReader(os.Stdin)
	// 

	var score int

	fmt.Println("Press enter to start the quiz")
	consoleReader.ReadString('\n')

	// Initiate Timer
	time.AfterFunc(time.Second * time.Duration(*seconds), func() {
		fmt.Println("Times up!")
	  fmt.Printf("You got %d out of %d questions correct.", score, len(questions))
		os.Exit(1)
	})
	// 

	// Iterate through each question asking questions and updating score
	fmt.Println("Please answer the following questions:")
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