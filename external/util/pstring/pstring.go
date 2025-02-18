package pstring

import (
	"strconv"
)

func GetInt(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return v
}
