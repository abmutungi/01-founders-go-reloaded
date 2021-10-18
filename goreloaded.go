package main

import ("fmt"
		"os"
		"strings"
)
func readFile() string{
	fileOne, err := os.ReadFile(os.Args[1])
	if err != nil{
		fmt.Println("Invalid file. Please try again")
}
return string(fileOne)
}

func toHex(s string) string{
	sliceOfS := strings.Split(s, " ")
	var emptySlice string

	for a := len(sliceOfS)-1 ; a >= 0 ; a--{
		if sliceofS[a] != "(hex)"{
			emptySlice += sliceOfS[a]
			fmt.Print(emptySlice)
		} else if sliceOfS[a] == "(hex)" {
			a -= 1
			sliceOfBytes := []byte(sliceOfS[a-1])
			encodedStr := hex.EncodeToString(sliceOfBytes)
		}
	}
}