package main

import (
	"fmt"
)

func displayEachSubjectScore(scores map[string]float64) {
	fmt.Println("------------------------------------------------------")
	fmt.Println("Subject Scores")
	fmt.Println("------------------------------------------------------")
	for key, value := range scores {
		fmt.Printf("%s: %.2f\n", key, value)
	}
	fmt.Println("------------------------------------------------------")
}

func displayResults(name string, average float64) {
	fmt.Println()
	fmt.Println("======================================================")
	fmt.Println("======================================================")
	fmt.Printf("Student Name: %s\n", name)
	fmt.Printf("Average Score: %.2f\n", average)
	fmt.Println("======================================================")
	fmt.Println("======================================================")
	fmt.Println()
}

func clear() {
	fmt.Print("\033[H\033[2J")
	introduction()
}
