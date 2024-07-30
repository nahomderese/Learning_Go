package main

func main() {

	clear()

	var name string = getStudentName()

	clear()

	var coursesLen int = getNumberOfCourses()

	clear()

	var scores map[string]float64 = getScoresForEachSubject(coursesLen)

	clear()

	var scoresArray []float64 = make([]float64, 0, len(scores))
	for _, value := range scores {
		scoresArray = append(scoresArray, value)
	}
	var average float64 = averageCalculator(scoresArray)
	displayResults(name, average)
	displayEachSubjectScore(scores)
}

func averageCalculator(scores []float64) float64 {
	var sum float64 = 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	return sum / float64(len(scores))
}
