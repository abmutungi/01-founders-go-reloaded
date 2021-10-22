package main

import (
	"fmt"
	"gr"
	"strings"
)

func main() {
	// Read the file and convert to slice of string
	sliceOfStr := strings.Split(gr.ReadFile(), " ")

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
			sliceOfStr[i] = gr.Atoi(sliceOfStr[i+1])
			sliceOfStr = append(sliceOfStr[:i], sliceOfStr[i+2:]...)
			// case "(up,":
			// case "(low,":

		}
	}
	fmt.Println(sliceOfStr)
}
