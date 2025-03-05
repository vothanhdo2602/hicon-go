package pstring

import (
	"fmt"
	"strconv"
)

func GetInt(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return v
}

func InterfaceToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", v)
	}
}
