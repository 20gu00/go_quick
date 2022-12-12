package common

import (
	"crypto/md5"
	"encoding/hex"
)

// md5加密用的secret
const secret = "github.com/20gu00"

// encrypt使用md5算法给password加密
func MD5(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword))) //16进制string
}
