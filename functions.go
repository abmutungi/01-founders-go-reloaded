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
	"strings"
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

func Atoi(s string) int {
	num := 0                // initalise num as 0
	for _, val := range s { // uses for range loop to iterate over s
		digit := 0
		if val <= '9' && val >= '0' {
			digit = int(val - '0')
			num = num*10 + digit
		}
	}
	return num
}

func removeSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
