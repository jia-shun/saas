package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Encode(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
