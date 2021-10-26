package main

import (
	"fmt"
	"gr"
	"strings"
)

func removeSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func main() {
	// Read the file and convert to slice of string
	sliceOfStr := strings.Split(gr.ReadFile(), " ")

	// for loops for multiple modifications of string

	for i := 0; i < len(sliceOfStr); i++ {
		n := 0

		if sliceOfStr[i] == "(cap," {
			n = gr.Atoi(sliceOfStr[i+1])
			for j := 1; j <= n; j++ {
				sliceOfStr[i-j] = strings.Title(sliceOfStr[i-j])
			}
		}
	}
	for i := 0; i < len(sliceOfStr); i++ {
		n := 0

		if sliceOfStr[i] == "(up," {
			n = gr.Atoi(sliceOfStr[i+1])
			for j := 1; j <= n; j++ {
				sliceOfStr[i-j] = strings.ToUpper(sliceOfStr[i-j])
			}
		}
	}
	for i := 0; i < len(sliceOfStr); i++ {
		n := 0

		if sliceOfStr[i] == "(low," {
			n = gr.Atoi(sliceOfStr[i+1])
			for j := 1; j <= n; j++ {
				sliceOfStr[i-j] = strings.ToLower(sliceOfStr[i-j])
			}
		}
	}

	// switch case statement for simple modifications of string

	for i, word := range sliceOfStr {
		switch word {
		case "(cap)":
			sliceOfStr[i-1] = strings.Title(strings.ToLower(sliceOfStr[i-1]))
			sliceOfStr = append(sliceOfStr[:i], sliceOfStr[i+1:]...)
		case "(low)":
			sliceOfStr[i-1] = strings.ToLower(sliceOfStr[i-1])
			sliceOfStr = append(sliceOfStr[:i], sliceOfStr[i+1:]...)
		case "(up)":
			sliceOfStr[i-1] = strings.ToUpper(sliceOfStr[i-1])
			sliceOfStr = append(sliceOfStr[:i], sliceOfStr[i+1:]...)
		case "(hex)":
			sliceOfStr[i-1] = gr.ToHex(sliceOfStr[i-1])
			sliceOfStr = append(sliceOfStr[:i], sliceOfStr[i+1:]...)
		case "(bin)":
			sliceOfStr[i-1] = gr.ToBin(sliceOfStr[i-1])
			sliceOfStr = append(sliceOfStr[:i], sliceOfStr[i+1:]...)
		case "a":
			sliceOfStr[i] = gr.ChangeA(sliceOfStr[i+1])
		case "A":
			sliceOfStr[i] = gr.ChangeA(sliceOfStr[i+1])
			sliceOfStr[i] = strings.Title(strings.ToLower(sliceOfStr[i]))
		case "(cap,":
			sliceOfStr = append(sliceOfStr[:i], sliceOfStr[i+2:]...)
		case "(low,":
			sliceOfStr = append(sliceOfStr[:i], sliceOfStr[i+2:]...)
		case "(up,":
			sliceOfStr = append(sliceOfStr[:i], sliceOfStr[i+2:]...)
		}
	}

	restring := strings.Join(sliceOfStr, " ")

	char := []rune(restring)

	for i := 0; i < len(char); i++ {
		if char[i] == ',' && char[i-1] == ' ' || char[i] == '.' && char[i-1] == ' ' || char[i] == '!' && char[i-1] == ' ' || char[i] == '?' && char[i-1] == ' ' || char[i] == ':' && char[i-1] == ' ' || char[i] == ';' && char[i-1] == ' ' { //|| char[i] == 39 && char[i-1] == 32 {
			char[i], char[i-1] = char[i-1], char[i]
		}
	}

	count := 0

	for i := 0; i < len(char); i++ {
		if char[i] == 39 {
			count++
			if char[i] == 39 && count%2 != 0 && char[i+1] == 32 {
				char[i], char[i+1] = char[i+1], char[i]
				i++
				fmt.Println("Hello")
			}
			if char[i] == 39 && count%2 == 0 && char[i-1] == 32 {
				char[i], char[i-1] = char[i-1], char[i]
				char[i-2], char[i-1] = char[i-1], char[i-2]
				fmt.Println("Bye")
			}
		}
	}
	a := removeSpace(string(char))
	fmt.Println(a)
}
