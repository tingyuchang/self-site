package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodingWithMD5(input string) string {
	m := md5.New()
	m.Write([]byte(input))

	return hex.EncodeToString(m.Sum(nil))
}