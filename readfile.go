package gr

import (
	"fmt"
	"os"
)

func ReadFile() string {
	fileOne, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Invalid file. Please try again")
	}
	return string(fileOne)
}
