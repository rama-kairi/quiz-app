package main

import (
	"fmt"

	"github.com/rama-kairi/quiz-app/data"
	"github.com/rama-kairi/quiz-app/helpers"
)

func main() {
	score := 0
	quizData := data.DataGenerator()

	// Initiate the quiz
	fmt.Println("Welcome to the Quiz App!")

	fmt.Println("Please select a topic of your choice: ")

	// Get all the Unique Topics
	topics := data.GetUniqueTopic(quizData)

	// Print the topics
	for i, t := range topics {
		fmt.Printf("%d. %s\n", i+1, t)
	}

	// Get the user input
	inputNumTopics := helpers.GetValidInput(len(topics))

	// Get the topic from topics with the index
	userTopic := topics[inputNumTopics-1]

	// Filter the quizData based on the topic
	quizDataWithTopic := data.FilterQuizData(quizData, userTopic)

	// Shuffle the questions
	quizDataWithTopic = helpers.ShuffleMap(quizDataWithTopic)

	// Shuffle the answer options
	for i, q := range quizDataWithTopic {
		answer_options := q["answer_options"].([]string)
		quizDataWithTopic[i]["answer_options"] = helpers.ShuffleArray(answer_options)
	}

	// Ask the user to select the number of questions
	fmt.Printf("How many questions do you want to answer? (Max: %d): ", len(quizDataWithTopic))
	inputNumQuestions := helpers.GetValidInput(len(quizDataWithTopic))

	// Slice the quizDataWithTopic based on the number of questions
	quizData = quizDataWithTopic[:inputNumQuestions]

	for i, q := range quizData {
		question := q["question"]
		answer_options := q["answer_options"].([]string)

		// Print the question
		fmt.Printf("Q No %d: %s\n", i+1, question)

		// Print the answer options
		for j, a := range answer_options {
			fmt.Printf("%d. %s\n", j+1, a)
		}

		// Take the user input
		inputNum := helpers.GetValidInput(len(answer_options))

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

	// Print the score
	fmt.Printf("Your score is %d out of %d\n", score, len(quizData))
}
