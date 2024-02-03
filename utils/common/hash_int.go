package common

import (
	"crypto/sha512"
	"encoding/hex"
	"strconv"
)

const sumContent int64 = 273387464515

func HashInt64(a int64) string {
	hash := sha512.New()
	hash.Write([]byte(strconv.FormatInt(a, 36)))
	hashed := hash.Sum([]byte(strconv.FormatInt(sumContent, 36)))
	return hex.EncodeToString(hashed)
}
