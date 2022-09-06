package auth

import (
	//"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"

	//"encoding/hex"
	///"errors"
	//"fmt"
	"io"

	"github.com/gps/pkg/pwd"
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

func encryptGCM(cipherBlock cipher.Block, plaintext string) (string, error) {
	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        return "", err
    }

	plaintextBytes := []byte(plaintext)
	ciphertext := gcm.Seal(nil, nonce, plaintextBytes, nil)
	return exportCipherText(ciphertext, nonce, gcm.Overhead())

}

func exportCipherText(ciphertext []byte, nonce []byte, size int) (string, error) {
	nonceSize := len(nonce)
	cipherLength := len(ciphertext) + nonceSize + 2
	
	data := make([]byte, cipherLength)
	data[0] = byte(nonceSize)

	idx := 1
	copy(data[idx:], nonce[0:nonceSize])

	idx += nonceSize

	copy(data[idx:], ciphertext)

	return base64.StdEncoding.EncodeToString(data), nil

}

/*
Encrypt text with secret key
*/
func Encrypt(plaintext string, secretkey string) (string, error) {
	secretkey = normalizeSecreyKey(secretkey)

	key := []byte(secretkey)
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	return encryptGCM(cipherBlock, plaintext)
	
}

func normalizeSecreyKey(secretkey string) string {
	if len(secretkey) < 32 { // needs to be 32 bytes
		remainingBytes := 32 - len(secretkey)
		bytesToFill := fillBytes(remainingBytes) // fill extra bytes
		secretkey += bytesToFill
	} else if len(secretkey) > 32 { // larger than 32 bytes
		secretkey = secretkey[0:32] // slice string to 32 bytes
	}
	return secretkey
}

func importCipherText(data []byte) ([]byte, []byte, int) {
	nonceSize := int(data[0])
	idx := 2
	size := int(data[idx])
	nonce, encryptedBytes := data[idx:idx+nonceSize], data[idx+nonceSize:]
	return encryptedBytes, nonce, size

}

func decryptGCM(cipherBlock cipher.Block, encryptedBytes[]byte, nonce []byte, size int) (string, error) {
	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", err
	}

	decryptedBytes, err := gcm.Open(nil, nonce, encryptedBytes, nil)
	if err != nil {
		return "", err
	}

	return string(decryptedBytes[:]), nil
}

/*
Decrypt text with key
*/
func Decrypt(ciphertext string, secretKey string) (string, error) {
	secretKey = normalizeSecreyKey(secretKey)
	key := []byte(secretKey)

	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	encryptedBytes, nonce, size := importCipherText(data)
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	return decryptGCM(cipherBlock, encryptedBytes, nonce, size)

}

func textToBinary(text string) string {
	bin := ""
	for _, c := range text {
		binary := pwd.TextToBinary(c)
		bin += binary
	}

	return bin
}

func xor(a rune, b rune) rune {
	if a == b {
		return '0'
	}
	return '1'
}

func groupBinary(binaryString string, limit int) []string {
	res := make([]string, 0)
	for i := 0; i < len(binaryString); i+=limit {
		line := binaryString[i:i+limit]
		res = append(res, line)
	}
	return res
}

func xorString(textA string, textB string) string {
	if len(textA) != len(textB) {
		panic("string not equal!")
	}
	str := ""
	for i := 0; i < len(textA); i++ {
		xorRes := xor(rune(textA[i]), rune(textB[i]))
		str += string(xorRes)
	}
	return str
}

const UNDERSCORE_BINARY = "01011111"
const SEMICOLON_BINARY  = "00111011"

func EncryptXOR(key string, plaintext string) {
	keyBinary := textToBinary(key)
	plaintextBinary := textToBinary(plaintext)

	keyBinaryGroup := groupBinary(keyBinary, 8)
	plaintextBinaryGroup := groupBinary(plaintextBinary, 8)

	keyBinaryGroupLength := len(keyBinaryGroup)

	finalResult := ""

	for i := 0; i < len(plaintextBinaryGroup); i++ {
		keyPos := i % keyBinaryGroupLength // return back to index to if finish the loop
		// only for key, not plaintext!

		if plaintextBinaryGroup[i] == UNDERSCORE_BINARY {
			finalResult += plaintextBinaryGroup[i] // skip underscore binary in plaintext
			continue
		}
		xorResult := xorString(plaintextBinaryGroup[i], keyBinaryGroup[keyPos])

		finalResult += xorResult
	}

	


}
