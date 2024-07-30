package internal

func AverageCalculator(scores map[string]float64) float64 {

	if len(scores) == 0 {
		return 0
	}

	scoresArray := make([]float64, 0, len(scores))
	for _, value := range scores {
		scoresArray = append(scoresArray, value)
	}
	var sum float64 = 0
	for i := 0; i < len(scoresArray); i++ {
		sum += scoresArray[i]
	}
	return sum / float64(len(scores))
}
