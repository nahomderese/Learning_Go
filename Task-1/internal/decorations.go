package internal

import (
	"fmt"
)

func DisplayEachSubjectScore(scores map[string]float64) {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("------------------------------------------------------")
	fmt.Println("Subject Scores")
	fmt.Println("------------------------------------------------------")
	for key, value := range scores {
		// format like a table
		fmt.Printf("%-20s %.2f\n", key, value)
		// fmt.Printf("%s: %.2f\n", key, value)
	}
	fmt.Println("------------------------------------------------------")
}

func DisplayResults(name string, average float64) {
	fmt.Println()
	fmt.Println()
	fmt.Println("======================================================")
	fmt.Println("======================================================")
	fmt.Printf("Student Name: %s\n", name)
	fmt.Printf("Average Score: %.2f\n", average)
	fmt.Println("======================================================")
	fmt.Println("======================================================")
	fmt.Println()
}

func Clear() {
	fmt.Print("\033[H\033[2J")
	Introduction()
}

func Introduction() {
	fmt.Println("------------------------------------------------------")
	fmt.Println("------------------------------------------------------")
	fmt.Println("This program calculates the average of student grades.")
	fmt.Println("------------------------------------------------------")
	fmt.Println("------------------------------------------------------")
}
