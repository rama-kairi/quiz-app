package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rama-kairi/quiz-app/data"
)

// We need json data for quiz

// We'll as questions from a json file ro the user in the terminal

// Then we'll list all the answer options

// User will enter the answer

// We'll check if the answer is correct or not

// We'll keep track of the score

// At the end we'll show the score

func main() {
	score := 0

	quizData := data.DataGenerator()

	// TODO: Ask the user to select a topic
	// TODO: Filter the quizData based on the topic
	// TODO: Shuffle the questions
	// TODO: Shuffle the answer options
	// TODO: Ask the user to select the number of questions

	for i, q := range quizData {
		question := q["question"]
		answer_options := q["answer_options"].([]string)
		// correct_answer := q["correct_answer"]

		// Print the question
		fmt.Printf("Q No %d: %s\n", i+1, question)

		// Print the answer options
		for j, a := range answer_options {
			fmt.Printf("%d. %s\n", j+1, a)
		}

		// Take the user input
		fmt.Print("Enter your answer: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		input = strings.TrimSpace(input)

		// Convert input to integer
		inputNum, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}

		// TODO: How to handle invalid input and keep asking for input until
		// the user enters a valid input and  keep on the same question

		if inputNum < 1 || inputNum > len(answer_options) {
			fmt.Println("Invalid input. Please enter a number between 1 and", len(answer_options))
			continue
		}

		// Get the correct answer from answer_options with the index
		correct_answer := answer_options[inputNum-1]

		// Check if the answer is correct
		if correct_answer == q["correct_answer"] {
			fmt.Println("Correct Answer!")
			score++
		} else {
			fmt.Println("Wrong Answer!")
		}
	}
}
