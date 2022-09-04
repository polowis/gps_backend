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