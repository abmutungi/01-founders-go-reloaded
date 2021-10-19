package gr

import (
	"strconv"
	"strings"
)

func ToHex(s string) string {
	sliceOfStr := strings.Split(s, " ")
	var emptySlice string

	for a := len(sliceOfStr) - 1; a >= 0; a-- {
		if sliceOfStr[a] != "(hex)" {
			emptySlice = emptySlice + " " + sliceOfStr[a]
		} else if sliceOfStr[a] == "(hex)" {
			a -= 1
			var replace string
			numberStr := strings.Replace(sliceOfStr[a], replace, "", -1)
			numberStr = strings.Replace(numberStr, replace, "", -1)

			output, _ := strconv.ParseInt(numberStr, 16, 64)
			inString := strconv.FormatInt(output, 10)
			emptySlice = emptySlice + " " + inString
		}
	}
	return emptySlice
}

func ToBin(s string) string {
	sliceOfStr := strings.Split(s, " ")
	var emptySlice string

	for a := len(sliceOfStr) - 1; a >= 0; a-- {
		if sliceOfStr[a] != "(bin)" {
			emptySlice = emptySlice + " " + sliceOfStr[a]
		} else if sliceOfStr[a] == "(bin)" {
			a -= 1
			var replace string
			numberStr := strings.Replace(sliceOfStr[a], replace, "", -1)
			numberStr = strings.Replace(numberStr, replace, "", -1)

			output, _ := strconv.ParseInt(numberStr, 2, 64)
			inString := strconv.FormatInt(output, 10)
			emptySlice = emptySlice + " " + inString
		}
	}
	return emptySlice
}

// func ToUpper(s string) string {
// 	sliceOfStr := strings.Split(s, " ")
// 	var emptySlice string

// 	for a := len(sliceOfStr) - 1; a >= 0; a-- {
// 		if sliceOfStr[a] != "(cap)" {
// 			emptySlice = emptySlice + " " + sliceOfStr[a]
// 		} else if sliceOfStr[a] == "(cap)" {
// 			a -= 1
// 			sliceOfStr[a] = strings.Title(sliceOfStr[a])
// 			sliceOfStr = append(sliceOfStr[:a], sliceOfStr[a+1:]...)
// 			emptySlice = emptySlice + " " + sliceOfStr[a]
// 		}
// 	}
// 	return emptySlice
// }

// func ToUpper(s string) string {
// 	sliceOfStr := strings.Split(s, " ")

// 	for i, word := range sliceOfStr {
// 		if word == "(cap)" {
// 			sliceOfStr[i-1] = strings.ToTitle(sliceOfStr[i-1])
// 			sliceOfStr = append(sliceOfStr[:i], sliceOfStr[i+1:]...)
// 		}
// 	}

// }
