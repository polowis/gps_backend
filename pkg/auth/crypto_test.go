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

func TestAESEncrypt35byteKey(t *testing.T) {
	key := "my super long secrey key for AES 35"
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

func TestAESDecryptShortKey(t *testing.T) {
	key := "my super long secrey key for AES"
	secret := "this text to be hide"
	text, _ := Encrypt(secret, key)
	fmt.Println(text)
	decryptedText, err := Decrypt(text, key)
	
	if err != nil {
		panic(err)
	}

	fmt.Println(decryptedText)
}

func TestTextToBinaryWithUnderscore(t *testing.T) {
	str := textToBinary("1_2")
	expected := "001100010101111100110010"
	if str != expected {
		t.Fatalf("result not correct")
	}
}

func TestTextToBinaryWithSemiColon(t *testing.T) {
	data := "12_3;"
	str := textToBinary(data)
	expected := "0011000100110010010111110011001100111011"
	if str != expected {
		t.Fatalf("result not correct")
	}
}