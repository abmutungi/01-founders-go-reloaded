// split

// if else or switch case
// hex, bin, cap, low
// (call upon created functions)

// separate functions for punc, change a/an

package gr

import (
	"fmt"
	"os"
	"regexp"
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
	return fmt.Sprint(output)
}

func ToBin(s string) string {
	output, _ := strconv.ParseInt(s, 2, 64)
	return fmt.Sprint(output)
}

// func punc ()

func ChangeA(s string) string {
	matched, _ := regexp.MatchString(`^[aeiouh]`, s)

	if matched == true {
		return "an"
	}
	return "a"
}

//"<number>)"

func TrimAtoi(s string) int {
	num := 0                // initalise num as 0
	result := 1             // initialise result as 1
	for i, val := range s { // uses for range loop to iterate over s
		digit := int(val) - 48
		if digit <= 9 && digit >= 0 {
			num = num*10 + digit
		} else if digit == -3 && i == 0 {
			result = -1
		} else if digit == -5 && i == 0 {
			result = 1
		} else {
			return 0
		}
	}
	num *= result
	return num
}

func HexNumber(s []string) []string {
	pd := 0
	var toAdd []string

	for range s {
	}
}
