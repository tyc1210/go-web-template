package util

import (
	"strconv"
)

func StrToInt(s string) (int, error) {
	num, err := strconv.Atoi(s)
	return num, err
}

func StrToUInt32(s string) (uint, error) {
	num, err := strconv.Atoi(s)
	return uint(num), err
}
