package texture

import (
	"math/rand"
	"testing"
	"time"
)

var width  = 16
var height = 16

func TestNewTexture(t *testing.T) {
	text := NewTexture(width, height)
	text.Save("1", "../../storage/sp")
	rand.Seed(time.Now().UnixNano())
}

func TestSliceSecretKey(t *testing.T) {
	key := sliceKey("abcdefghjkl")
	if key != "abcdefg" {
		t.Fatalf("Key not match")
	}
}

func TestTextToBinary(t *testing.T) {
	text := "abcd"
	binary := textToBinary(text)
	expected := "01100001011000100110001101100100"
	if binary != expected {
		t.Fatalf("Binary result not match")
	}
}

func TestTextureId(t *testing.T) {
	text := NewTexture(width, height)

	text.SetKey("abcdefghjkl")
	text.SetCode("13gdh31leq13ll098dmmzkd")
	expected := "505104000d5556"
	if expected != text.ID() {
		t.Fatalf("Hex result not match")
	}
}
