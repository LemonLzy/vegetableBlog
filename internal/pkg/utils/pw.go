package utils

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
	"unsafe"
)

// scryptPw 使用scrypt算法加密
func scryptPw(password string) string {
	key, _ := scrypt.Key([]byte(password), []byte{'a', 'b', 'c'}, 1<<15, 8, 1, 32)
	return base64.StdEncoding.EncodeToString(key)
}

// BcryptPw 使用bcrypt算法加密
func BcryptPw(password string) string {
	fromPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return Bytes2String(fromPassword)
}

// BcryptCompare 判断密码是否相等
func BcryptCompare(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

// Bytes2String 字节数组强转字符串
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
