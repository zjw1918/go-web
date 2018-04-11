package utils

import "reflect"

//ConvertToInt64 ...
func ConvertToInt64(number interface{}) int64 {
	if reflect.TypeOf(number).String() == "uint" {
		return int64(number.(uint))
	}
	return number.(int64)
}
