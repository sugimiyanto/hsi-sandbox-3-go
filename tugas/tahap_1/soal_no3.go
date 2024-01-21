package main

import (
	"fmt"
	"errors"
)

func formatError(message string) error {
	return fmt.Errorf("return from formatError: %s", message)
}


func main() {
	errMessage := "error 1"
	formattedError := fmt.Errorf("Output Errorf: %s", errMessage)
	fmt.Println(formattedError)

	err := formatError(errMessage)
	fmt.Printf("%T\n", err)

	// errors.New("Output errors.New: %s", err)
	// ini akan error seperti di bawah kalau dijalankan:
		/*
		# command-line-arguments
		./soal_no3.go:12:38: too many arguments in call to errors.New
        have (string, string)
        want (string)
		*/
	errMessageNew := errors.New("Output errors.New")
	fmt.Println(errMessageNew)
	fmt.Printf("%T", errMessageNew)
}