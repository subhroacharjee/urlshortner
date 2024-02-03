package common

import "strconv"

func ParseUint(str string) (*uint, error) {
	val, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return nil, err
	}
	res := uint(val)
	return &res, err
}
