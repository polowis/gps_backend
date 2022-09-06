package auth

import (
	"fmt"
	"testing"
)

func TestAESEncrypt32byteKey(t *testing.T) {
	key := "my super long secrey key for AES"
	secret := "this text to be hide"
	text, err := Encrypt(secret, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}

func TestAESEncrypt33byteKey(t *testing.T) {
	key := "my super long secrey key for AES 33"
	secret := "this text to be hide"
	text, err := Encrypt(secret, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}

func TestAESEncryptShortKey(t *testing.T) {
	key := "0123456"
	secret := "this text to be hide"
	text, err := Encrypt(secret, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}