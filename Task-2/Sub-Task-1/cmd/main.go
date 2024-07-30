package main

import (
	"fmt"

	"github.com/Nahom-Derese/Learning_Go/Task-2/Sub-Task-1/internal"
)

func main() {
	fmt.Println(internal.WordCounter("Hello World"))
	fmt.Println(internal.WordCounter("Hello World, world"))
	fmt.Println(internal.WordCounter("Hello hello"))
	fmt.Println(internal.WordCounter("Hello hello, world"))
}
