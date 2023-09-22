package tools

import (
	"crypto/md5"
	"encoding/hex"
)

func GenMD5(s string) (md5String string) {
	h := md5.New()
	h.Write([]byte(s))
	md5String = hex.EncodeToString(h.Sum(nil))
	return
}
