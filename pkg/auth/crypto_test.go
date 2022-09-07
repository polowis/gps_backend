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

func TestEncryptSimpleXOR(t *testing.T) {
	plaintext := "20_35;1_2"
	key := "134567"
	expected := "000000110000001101011111000001100000001100111011000000000101111100000110"
	ciphertext := EncryptXOR(plaintext, key)
	if ciphertext != expected {
		t.Fatalf("result not match!")
	}
}

func TestEncryptSimpleXORWithBase64(t *testing.T) {
	plaintext := "20_35;1_2"
	key := "134567"
	expected := "000000110000001101011111000001100000001100111011000000000101111100000110"
	ciphertext := EncryptXOR(plaintext, key)
	if ciphertext != expected {
		t.Fatalf("xor result not match!")
	}
	base64Expected := "MDAwMDAwMTEwMDAwMDAxMTAxMDExMTExMDAwMDAxMTAwMDAwMDAxMTAwMTExMDExMDAwMDAwMDAwMTAxMTExMTAwMDAwMTEw"
	base64cipher := Base64Encode(ciphertext)
	if base64cipher != base64Expected {
		t.Fatalf("Base64 result not match")
	}
}


func TestEncryptXORWithHex(t *testing.T) {
	plaintext := "20_35;1_2"
	key := "134567"
	expected := "000000110000001101011111000001100000001100111011000000000101111100000110"
	ciphertext := EncryptXOR(plaintext, key)
	if ciphertext != expected {
		t.Fatalf("xor result not match!")
	}
	expectedHex := "335f633b05f6"
	hexcipher, err := BinaryToHex(ciphertext)
	if err != nil {
		t.Fatal(err)
	}
	
	if hexcipher != expectedHex {
		t.Fatalf("hex result not matched!")
	}

}

func TestDecryptSimpleXOR(t *testing.T) {
	plaintext := "20_35;1_2"
	key := "134567"
	plaintextBinary := textToBinary(plaintext)
	expected := "000000110000001101011111000001100000001100111011000000000101111100000110"
	ciphertext := EncryptXOR(plaintext, key)
	if ciphertext != expected {
		t.Fatalf("xor result not match!")
	}

	output := DecryptXOR(ciphertext, key)
	if output != plaintextBinary {
		t.Fatalf("unable to decrypt")
	}
}