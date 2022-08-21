package util

import (
	"crypto/md5"
	"encoding/hex"
)

//hash value 128位元16進位32字元字串, 用於確保訊息傳輸完整一致
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
