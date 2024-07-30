package main

import "github.com/Nahom-Derese/Learning_Go/Task-1/internal"

func main() {

	internal.Clear()

	name := internal.GetStudentName()

	internal.Clear()

	coursesLen := internal.GetNumberOfCourses()

	internal.Clear()

	scores := internal.GetScoresForEachSubject(coursesLen)

	internal.Clear()

	average := internal.AverageCalculator(scores)

	internal.DisplayEachSubjectScore(scores)
	internal.DisplayResults(name, average)
}
