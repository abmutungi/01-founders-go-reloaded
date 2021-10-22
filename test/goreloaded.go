package main

import (
	"fmt"
	"gr"
	"strings"
)

func main() {
	// Read the file and convert to slice of string
	textFile := gr.ReadFile()
	sliceOfStr := strings.Split(textFile, " ")

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
			// case "(cap,":
			// 	break
			// sliceOfStr = append(sliceOfStr[:i], sliceOfStr[i+2:]...)
			// case "(up,":
			// case "(low,":

		}
	}

	for i := 0; i < len(sliceOfStr); i++ {
		n := 0
		var newSlice []string

		if sliceOfStr[i] == "(cap," {
			n = gr.Atoi(sliceOfStr[i+1])
			fmt.Println(n)
			for j := 1; j <= n; j++ {
				sliceOfStr[i-j] = strings.Title(sliceOfStr[i-j])
				newSlice = append(newSlice, sliceOfStr[i-j])
			}
		}
	}
	fmt.Println(sliceOfStr)
}
