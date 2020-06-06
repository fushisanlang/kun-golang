package until

import (
	"fmt"
	"strconv"
)

func TypeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
func StrToInt(StrIn string) int {
	IntOut, err := strconv.Atoi(StrIn)
	CheckErr(err)
	return IntOut
}

