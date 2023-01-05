package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

// GetValidInput - Get valid input from the user
func GetValidInput(maxNum int) int {
	for {
		// Take the user input
		fmt.Printf("Enter a number between 1 and %d: ", maxNum)
		var inputNum int
		_, err := fmt.Scanf("%d", &inputNum)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		if inputNum < 1 || inputNum > maxNum {
			fmt.Println("Invalid input. Please enter a number between 1 and", maxNum)
		} else {
			return inputNum
		}
	}
}

// ShuffleMap - Shuffle the Map
func ShuffleMap(arr []map[string]interface{}) []map[string]interface{} {
	for i := range arr {
		rand.Seed(time.Now().UnixNano())
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// ShuffleArray - Shuffle the Array
func ShuffleArray(arr []string) []string {
	for i := range arr {
		rand.Seed(time.Now().UnixNano())
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
