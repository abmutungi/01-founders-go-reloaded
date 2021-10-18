package gr

import (
	"fmt"
	"strconv"
	"strings"
)

func ToHex(s string) string {
	sliceOfStr := strings.Split(s, " ")
	var emptySlice string

	for a := len(sliceOfS) - 1; a >= 0; a-- {
		if sliceOfStr[a] != "(hex)" {
			emptySlice += sliceOfStr[a]
			fmt.Print(emptySlice)
		} else if sliceOfStr[a] == "(hex)" {
			a -= 1
			numberStr := strings.Replace(sliceOfStr[a], -1)
			numberStr = strings.Replace(numberStr, -1)

			output, _ := strconv.ParseInt(numberStr, 64)
			inString := strconv.FormatInt(output, 10)
			emptySlice = inString + " " + emptySlice
		}
	}
	return emptySlice
}

/*func toBin(s string) string {

}
*/
