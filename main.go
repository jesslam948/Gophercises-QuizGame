// Description: Gophercise 1 - Simple Quiz Game
// References:
//	https://courses.calhoun.io/lessons/les_goph_01
//	https://github.com/gophercises/quiz

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	// open the csv file
	questions, err := os.Open("questions.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer questions.Close()

	// parse the file
	reader := csv.NewReader(questions)

	// while loop going through each line
	for {
		line, err := reader.Read()
		if err == io.EOF {
			// reached end of file
			break
		}
		if err != nil {
			// there's an error
			fmt.Println(err)
		}

		// print out the line contents (question and answer)
		fmt.Printf("Question: %s Answer %s \n", line[0], line[1])
	}
}
