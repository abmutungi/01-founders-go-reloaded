package main

import (
	"fmt"
	"gr"
)

func main() {
	textfile := gr.ReadFile()
	convertHex := gr.ToHex(string(textfile))
	fmt.Println(convertHex)
}

/*read input file

story, _ := os.ReadFile(arguments[0])

text := strings.Split(string(story), " ")

for i, word := range text {
	switch word {
	case word == "(up)":
		text[i-1] = strings.ToUpper(text[i-1])
		text = append(text[:i], text[i+1:]...)

	case word == "(low)":
		text[i-1] = strings.ToLower(text[i-1])
		text = append(text[:i], text[i+1:]...)

	case word == "(cap)":
		text[i-1] = strings.ToTitle(text[i-1])
		text = append(text[:i], text[i+1:]...)

	default:
		fmt.Println(text)
	}
}
*/
