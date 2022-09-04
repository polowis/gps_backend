package pwd

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func generatePwd(letters string, length int) string {
	bytes := make([]byte, length)
    for i := range bytes {
        bytes[i] = letters[rand.Intn(len(letters))]
		rand.Seed(time.Now().UnixNano())
    }
    return string(bytes)
}

func GenerateNumber(length int) string {
	const letters = "0123456789"
	return generatePwd(letters, length)
}

func GenerateText(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    return generatePwd(letters, length)
}

func TextToBinary(text rune) string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "%08b", text)
    return buffer.String()
}