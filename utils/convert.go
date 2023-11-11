package utils

import "fmt"

func ConvertInterfaceToString(x interface{}) string {
	var str string
	if val, ok := x.(string); ok {
		str = val
	} else {
		str = fmt.Sprintf("%v", x)
	}
	return str
}
