package gr

import (
	"fmt"
	"strconv"
	"strings"
)

func ToHex(s string) string {
	sliceOfS := strings.Split(s, " ")
	var emptySlice string

	for a := len(sliceOfS) - 1; a >= 0; a-- {
		if sliceOfS[a] != "(hex)" {
			emptySlice += sliceOfS[a]
			fmt.Print(emptySlice)
		} else if sliceOfS[a] == "(hex)" {
			a -= 1
			numberStr := strings.Replace(s, "", -1)
			numberStr = strings.Replace(numberStr, "", -1)

			output, _ := strconv.ParseInt(toHex(), 16, 64)
		}
	}
}

/*func toBin(s string) string {

}
*/
