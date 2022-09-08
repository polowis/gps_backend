package auth

import (
	//"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strconv"

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
	return err == nil
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
const UNDERSCORE_HEX    = 95
const SEMICOLON_BINARY  = "00111011"
const SEMICOLON_HEX     = 59

/*
Plaintext to encrypt
Key - secret key

return BINARY
*/
func EncryptXOR(plaintext string, key string) string {
	keyBinary := textToBinary(key)
	plaintextBinary := textToBinary(plaintext)

	keyBinaryGroup := groupBinary(keyBinary, 8)
	plaintextBinaryGroup := groupBinary(plaintextBinary, 8)

	keyBinaryGroupLength := len(keyBinaryGroup)

	finalResult := ""

	for i := 0; i < len(plaintextBinaryGroup); i++ {
		keyPos := i % keyBinaryGroupLength // return back to index to if finish the loop
		// only for key, not plaintext!

		if hasSkippedCharacter(plaintextBinaryGroup[i]) {
			finalResult += plaintextBinaryGroup[i] // skip underscore binary in plaintext
			continue
		}
		xorResult := xorString(plaintextBinaryGroup[i], keyBinaryGroup[keyPos])

		finalResult += xorResult
	}
	return finalResult

}

/*
Ciphertext must be in binary format
*/
func DecryptXOR(cipherBinary string, key string) string {
	keyBinary := textToBinary(key)
	keyBinaryGroup := groupBinary(keyBinary, 8)
	cipherBinaryGroup := groupBinary(cipherBinary, 8)
	keyBinaryGroupLength := len(keyBinaryGroup)

	finalResult := ""

	for i := 0; i < len(cipherBinaryGroup); i++ {
		keyPos := i % keyBinaryGroupLength // return back to index to if finish the loop
		// only for key, not plaintext!
		var xorResult string
		decimal, _ := BinaryToDecimal(cipherBinaryGroup[i])
		if decimal == UNDERSCORE_HEX {  // DO NOT XOR
			// if decimal is underscore or semicolon
			xorResult = "_"
		} else if decimal == SEMICOLON_HEX {
			xorResult = ";"
		} else { // do xor calc if not underscore or
			xorResult = xorString(cipherBinaryGroup[i], keyBinaryGroup[keyPos])
			decimal, _ = BinaryToDecimal(xorResult) // convert to decimal after xor to get original number
			xorResult = string(rune(decimal))
			if _, err := strconv.Atoi(xorResult); err != nil {
				// if result is not a decimal convertible
				xorResult = strconv.FormatInt(decimal, 10) // fall back to only number
			} 
		}

		finalResult += xorResult
	}
	return finalResult
}

func hasSkippedCharacter(binary string) bool {
	return binary == UNDERSCORE_BINARY || binary == SEMICOLON_BINARY
}

func Base64Encode(text string) (response string){
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func Base64Decode(text string) (response string) {
	res, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "" // cant decode, fall back to default ""
	}
	return string(res)
}

func BinaryToHex(text string) (string, error) {
	binArray := groupBinary(text, 8)
	result := ""
	for _, bin := range binArray {
		ui, err := strconv.ParseUint(bin, 2, 64)
		if err != nil {
			return "", err
		}
		hexcipher := fmt.Sprintf("%x", ui)
		if len(hexcipher) == 1 { 
			// only 1 number avilable, ideally we want 2 so FF can become 255 (8bit)
			result += "0" // add leading 0 in bit
		}
		result += hexcipher
	}

    return result, nil
}

func BinaryToDecimal(bin string) (int64, error) {
	if i, err := strconv.ParseInt(bin, 2, 64); err != nil {
		return 0, err
	} else {
		return i, nil
	}
}

/*
Convert hexadecimal to binary
*/
func HexToBinary(text string) (string, error) {
	i, err := strconv.ParseUint(text, 16, 32)
	if err != nil {
		return "", err
	}
	bin := fmt.Sprintf("%08b", i)
	return bin, nil
}
