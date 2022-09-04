package pwd

import (
	"testing"
)

func TestTextToBinaryLowerCase(t *testing.T) {
	res := TextToBinary('a')
	if res != "01100001" {
		t.Fatalf("string not equal")
	}
}

func TestTextToBinaryUpperCase(t *testing.T) {
	res := TextToBinary('A')
	if res != "01000001" {
		t.Fatalf("string not equal")
	}
}

func TestBinaryToText(t *testing.T) {
	bin := "01100001011000100110001101100100"
	res, _ := BinaryToHex(bin)
	expected := "61626364"
	if res != expected {
		t.Fatalf("String not equal, expected %s but got %s", expected, res)
	}
}