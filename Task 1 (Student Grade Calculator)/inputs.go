package main

import (
	"fmt"
)

func introduction() {
	fmt.Println("------------------------------------------------------")
	fmt.Println("------------------------------------------------------")
	fmt.Println("This program calculates the average of student grades.")
	fmt.Println("------------------------------------------------------")
	fmt.Println("------------------------------------------------------")
}

func getStudentName() string {
	var name string
	fmt.Print("Enter student name: ")
	fmt.Scanln(&name)
	return name
}

func getNumberOfCourses() int {
	var numCourses int
	fmt.Print("Enter number of courses: ")
	fmt.Scanln(&numCourses)
	return numCourses
}

func getScoresForEachSubject(numCourses int) map[string]float64 {
	scores := make(map[string]float64)
	for i := 0; i < numCourses; i++ {
		var courseName string
		var score float64
		fmt.Println(" ")
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>x>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Print("Enter course name: ")
		fmt.Scanln(&courseName)
		fmt.Print("Enter score: ")
		fmt.Scanln(&score)
		scores[courseName] = score
	}
	return scores
}
