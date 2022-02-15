package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str string) string {
	md5Bit := md5.Sum([]byte(str))
	return hex.EncodeToString(md5Bit[:])
}
