package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func intInput(prompt string) int {
	fmt.Println()
	fmt.Println()
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	num, err := strconv.Atoi(strings.TrimSpace(input))

	for err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		fmt.Println()
		fmt.Println()
		fmt.Print(prompt)
		input, _ = reader.ReadString('\n')
		num, err = strconv.Atoi(strings.TrimSpace(input))
	}

	return num
}

func floatInput(prompt string) float64 {
	fmt.Println()
	fmt.Println()
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	for err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		fmt.Println()
		fmt.Println()

		fmt.Print(prompt)
		input, _ = reader.ReadString('\n')

		num, err = strconv.ParseFloat(strings.TrimSpace(input), 64)
	}

	return num
}

func stringInput(prompt string) string {
	fmt.Println()
	fmt.Println()
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')

	for len(input) < 2 {
		fmt.Println()
		fmt.Println()
		fmt.Println("Invalid input. Please enter a valid name.")
		fmt.Print(prompt)
		input, _ = reader.ReadString('\n')
	}

	return strings.TrimSpace(input)
}

func GetStudentName() string {
	return stringInput("Enter student name: ")
}

func GetNumberOfCourses() int {
	return intInput("Enter number of courses: ")
}

func GetScoresForEachSubject(numCourses int) map[string]float64 {
	scores := make(map[string]float64)
	for i := 0; i < numCourses; i++ {
		var courseName string
		fmt.Println(" ")
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>x>>>>>>>>>>>>>>>>>>>>>>")

		courseName = stringInput(fmt.Sprintf("Enter name of Course %d : ", i+1))
		numCourses := floatInput("Enter score: ")

		scores[courseName] = float64(numCourses)
	}
	return scores
}
