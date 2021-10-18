package main

import ("fmt"
		"os"
		"strings"
)

func main (){

	textfile := gr.readFile()
	convertHex:= toHex(string(textfile))
	fmt.Println(convertHex)
}