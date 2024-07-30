package main

import (
	"fmt"

	"github.com/Nahom-Derese/Learning_Go/Task-2/Sub-Task-2/internal"
)

func main() {
	fmt.Println(internal.IsPalindrome("Hello World"))
	fmt.Println(internal.IsPalindrome("Hello World, world"))
	fmt.Println(internal.IsPalindrome("Hello hello"))
	fmt.Println(internal.IsPalindrome("Hello hello, world"))
}
