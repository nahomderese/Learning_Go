package internal

import "strings"

func WordCounter(s string) map[string]int64 {
	var freq map[string]int64 = make(map[string]int64)

	var words []string = strings.Split(s, " ")

	for _, word := range words {
		cleaned := removeNonLetters(word)
		freq[strings.ToLower(cleaned)]++
	}

	return freq

}
