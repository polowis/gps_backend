package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
)

/*
Hash password with given plain text
*/
func Hash(plaintext string) string {
	password := []byte(plaintext)

    // default cost of 10
    hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
    return string(hashedPassword)

    
}

/*
Compare hashed password in database with plaintext password
*/
func HasHash(hashedPassword string, password string) bool {
	//Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil { // password not match
		return false
	}
	return true
}

/*
n = number of byte to fill
*/
func fillBytes(n int) string {
	str := ""
	for i := 0; i < n; i ++ {
		str += "0"
	}
	return str
}

/*
Encrypt text with secret key
*/
func Encrypt(plainText string, secretkey string) (*string, error) {
	if len(secretkey) < 32 { // needs to be 32 bytes
		remainingBytes := 32 - len(secretkey)
		bytesToFill := fillBytes(remainingBytes) // fill extra bytes
		secretkey += bytesToFill
	} else if len(secretkey) > 32 { // larger than 32 bytes
		secretkey = secretkey[0:32] // slice string to 32 bytes
	}

	plaintext := []byte(plainText)
	key := []byte(secretkey)
	

    cipherBlock, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    mode, err := cipher.NewGCM(cipherBlock)
    if err != nil {
        return nil, err
    }

    nonce := make([]byte, mode.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }

	cipherText := mode.Seal(nonce, nonce, plaintext, nil)

	res := fmt.Sprintf("%x", cipherText)

    return &res, nil
}

/*
Decrypt text with key
*/
func Decrypt(cipherText string, secretKey string) (*string, error) {
	ciphertext := []byte(cipherText)
	key := []byte(secretKey)

    cipherBlock, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    mode, err := cipher.NewGCM(cipherBlock)
    if err != nil {
        return nil, err
    }

    nonceSize := mode.NonceSize()
    if len(ciphertext) < nonceSize {
        return nil, errors.New("ciphertext too short")
    }

    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	byteText, err := mode.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	res := fmt.Sprintf("%s", byteText)
    return &res, nil
}
