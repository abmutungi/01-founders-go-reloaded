// split

// if else or switch case
// hex, bin, cap, low
// (call upon created functions)

// separate functions for punc, change a/an

package gr

import (
	"fmt"
	"os"
	"strconv"
)

func ReadFile() string {
	fileOne, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input")
	}
	return string(fileOne)
}

func ToHex(s string) string {

	output, _ := strconv.ParseInt(s, 16, 64)
	//inString := strconv.FormatInt(output, 10)
	return fmt.Sprint(output)

}

func ToBin(s string) string {

	output, _ := strconv.ParseInt(s, 2, 64)
	return fmt.Sprint(output)
}

//func punc ()
//change a/an
