package main

import (
	"fmt"
	"gr"
)

func main() {
	textfile := gr.ReadFile()
	convertHex := gr.ToHex(string(textfile))
	convertBin := gr.ToBin(convertHex)
	//	capitalise := gr.ToUpper(convertBin)
	fmt.Println(convertBin)
}
