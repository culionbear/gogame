package auth

import (
	"crypto/aes"
	"crypto/md5"
	"encoding/hex"
)

func (m *Manager) MD5(str string) string {
	Md5Handle := md5.New()
	Md5Handle.Write([]byte(str))
	return hex.EncodeToString(Md5Handle.Sum(nil))
}

func (m *Manager) AesEncrypt(buf []byte) string {
	cipher, _ := aes.NewCipher([]byte(m.config.Token.Key))
	length := (len(buf) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, buf)
	pad := byte(len(plain) - len(buf))
	for i := len(buf); i < len(plain); i++ {
		plain[i] = pad
	}
	res := make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(buf); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(res[bs:be], plain[bs:be])
	}
	return hex.EncodeToString(res)
}

func (m *Manager) AesDecrypt(token string) ([]byte, error) {
	buf, err := hex.DecodeString(token)
	if err != nil {
		return nil, err
	}
	return m.aesDecryptECB(buf), nil
}

func (m *Manager) aesDecryptECB(buf []byte) []byte {
	cipher, _ := aes.NewCipher([]byte(m.config.Token.Key))
	res := make([]byte, len(buf))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(buf); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(res[bs:be], buf[bs:be])
	}

	trim := 0
	if len(res) > 0 {
		trim = len(res) - int(res[len(res)-1])
	}

	return res[:trim]
}
