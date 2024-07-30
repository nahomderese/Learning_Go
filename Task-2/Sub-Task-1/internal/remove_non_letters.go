package internal

func removeNonLetters(input string) string {
	var result []rune

	for _, char := range input {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			result = append(result, char)
		}
	}
	return string(result)
}
