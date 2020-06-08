package until

import (
	"fmt"
	"strconv"
)

//判断属性
func TypeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

//str转int
func StrToInt(StrIn string) int {
	IntOut, err := strconv.Atoi(StrIn)
	CheckErr(err)
	return IntOut
}
