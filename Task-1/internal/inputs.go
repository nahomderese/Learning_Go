package internal

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func intInput(prompt string, reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	fmt.Println()
	fmt.Println()
	fmt.Print(prompt)
	scanner.Scan()
	input := scanner.Text()
	num, err := strconv.Atoi(strings.TrimSpace(input))

	for err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		fmt.Println()
		fmt.Println()
		fmt.Print(prompt)
		scanner.Scan()
		input = scanner.Text()
		num, err = strconv.Atoi(strings.TrimSpace(input))
	}

	return num
}

// func intInput(prompt string) int {
// 	fmt.Println()
// 	fmt.Println()
// 	fmt.Print(prompt)
// 	input, _ := reader.ReadString('\n')
// 	num, err := strconv.Atoi(strings.TrimSpace(input))

// 	for err != nil {
// 		fmt.Println("Invalid input. Please enter a valid number.")
// 		fmt.Println()
// 		fmt.Println()
// 		fmt.Print(prompt)
// 		input, _ = reader.ReadString('\n')
// 		num, err = strconv.Atoi(strings.TrimSpace(input))
// 	}

// 	return num
// }

func floatInput(prompt string, reader io.Reader) float64 {
	scanner := bufio.NewScanner(reader)
	fmt.Println()
	fmt.Println()
	fmt.Print(prompt)
	scanner.Scan()
	input := scanner.Text()
	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	for err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		fmt.Println()
		fmt.Println()

		fmt.Print(prompt)
		scanner.Scan()
		input := scanner.Text()

		num, err = strconv.ParseFloat(strings.TrimSpace(input), 64)
	}

	return num
}

func stringInput(prompt string, reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	fmt.Println()
	fmt.Println()
	fmt.Print(prompt)
	scanner.Scan()
	input := scanner.Text()

	for len(input) < 2 {
		fmt.Println()
		fmt.Println()
		fmt.Println("Invalid input. Please enter a valid name.")
		fmt.Print(prompt)
		scanner.Scan()
		input = scanner.Text()
	}

	return strings.TrimSpace(input)
}

func GetStudentName() string {
	return stringInput("Enter student name: ", reader)
}

func GetNumberOfCourses() int {
	return intInput("Enter number of courses: ", reader)
}

func GetScoresForEachSubject(numCourses int) map[string]float64 {
	scores := make(map[string]float64)
	for i := 0; i < numCourses; i++ {
		var courseName string
		fmt.Println(" ")
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>x>>>>>>>>>>>>>>>>>>>>>>")

		courseName = stringInput(fmt.Sprintf("Enter name of Course %d : ", i+1), reader)
		numCourses := floatInput("Enter score: ", reader)

		scores[courseName] = float64(numCourses)
	}
	return scores
}
