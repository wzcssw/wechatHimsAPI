package lib

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
