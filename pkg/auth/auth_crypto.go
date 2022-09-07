package auth


func HashOrder(order string) string {
	return Hash(order)
}

func HasMatchedOrder(order string, hashedText string) bool {
	return HasHash(hashedText, order)
}

/*
Encrypt coordinate, coordinate must be in plain text
Example 12_34; where undescore split x and y, semicolon split coordinates

@param plaintext coordinate must be in plain text

@param key, should be user box order

@Return base64 encoded cipher text
*/
func EncryptCoordinates(plaintext string, key string) string {
	xorCipher := EncryptXOR(plaintext, key)
	hexCipher, err := BinaryToHex(xorCipher)
	if err != nil {
		panic(err)
	}
	base64Cipher := Base64Encode(hexCipher)
	return base64Cipher
}

/*
@param ciphertext - ciphertext in base64 format
@param key - key that used in encryption method
@return original coordinate
*/
func DecryptCoordinates(ciphertext string, key string) string {
	base64Cipher := Base64Decode(ciphertext)
	binHexArray := groupBinary(base64Cipher, 2)
	result := ""
	for _, bin := range binHexArray {
		res, _ := HexToBinary(bin)
		result += res
	}

	//binaryCipher, err := HexToBinary(base64Cipher)
	//if err != nil {
	//	panic(err)
	//}
	return DecryptXOR(result, key)

}
