package internal

import "strings"

func IsPalindrome(s string) bool {

	s = removeNonLetters(s)
	s = strings.ToLower(s)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
