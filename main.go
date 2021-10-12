package main

import (
	"fmt"
	"strings"
	"os"
	"strconv"
)

func main() {
	
	arguments := os.Args[1:]

 	// error handling

	// if len(arguments) < 1 {
    //     fmt.Println("Please give one argument.")
	// }

	// if len(arguments) > 2 {
	// 	fmt.Println("Too many arguments")
	// }

	//read input file

	text,_ := os.ReadFile(arguments[0])

	word := strings.Split(text, " ")


	for i, word := range word {
		if word == "(up)"{
		word[i-1] = strings.ToUpper(i-1)
		}

	}