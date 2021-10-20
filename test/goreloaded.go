package main

import (
	"fmt"
	"gr"
	"strings"
)

func main() {
	//Read the file and convert to slice of string
	textfile := gr.ReadFile()

	sliceOfStr := strings.Split(textfile, " ")

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

		}
	}
	fmt.Println(sliceOfStr)
}
