package auth

import (
	"crypto/md5"
	"encoding/hex"
)

func (m *Manager)MD5(str string) string {
	Md5Handle := md5.New()
	Md5Handle.Write([]byte(str))
	return hex.EncodeToString(Md5Handle.Sum(nil))
}
