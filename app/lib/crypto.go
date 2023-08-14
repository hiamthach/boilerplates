package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func HashSHA256(key string, content string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(content))
	expectedMAC := mac.Sum(nil)
	return fmt.Sprintf("%x", expectedMAC)
}

var _IV = []byte("!Thach@hi0hihi")

func EncryptString(key string, content string) string {

	k1 := []byte(key)
	data := []byte(content)
	block, _ := aes.NewCipher(k1)
	stream := cipher.NewCFBEncrypter(block, _IV)
	stream.XORKeyStream(data, data)
	return fmt.Sprintf("%x", data)
}

func DecryptString(key string, content string) string {
	bytes, _ := hex.DecodeString(content)
	block, _ := aes.NewCipher([]byte(key))
	stream := cipher.NewCFBDecrypter(block, _IV)

	stream.XORKeyStream(bytes, bytes)

	return string(bytes)

}

func Base64String(content string) string {
	rawDecodedText := base64.StdEncoding.EncodeToString([]byte(content))
	return rawDecodedText
}

func DecodeBase64String(content string) string {
	rawDecodedText, _ := base64.StdEncoding.DecodeString(content)
	return string(rawDecodedText)
}

func GetMD5Hash(content string) string {
	hash := md5.Sum([]byte(content))
	return hex.EncodeToString(hash[:])
}
