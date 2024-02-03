package common

import "strconv"

func ParseInt(a string) (int, error) {
	return strconv.Atoi(a)
}

func ParseInt64(a string) (int64, error) {
	return strconv.ParseInt(a, 10, 64)
}
