package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	filename := "1 output of Println"
	number := 12345
	fmt.Println(filename, currentTime, number)
	currentTime = time.Now()
	filename = "2 output of Sprintln"
	number = 9876
	fmt.Sprintln(filename, currentTime, number)
	currentTime = time.Now()
	filename = "3 output of Sprintln"
	number = 9876
	formattedString := fmt.Sprintln(filename, currentTime, number)
	fmt.Print(formattedString)

	fmt.Printf("%T", formattedString)
}