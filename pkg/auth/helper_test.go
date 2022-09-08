package auth

import (
	"testing"
)

func TestConcatOrder(t *testing.T) {
	str := "12_13_23_24"
	expected := "12132324"
	res := concatOrder(str)
	if res != expected {
		t.Fatalf("not match")
	}
}

func TestNormalizeCoordinate(t *testing.T) {
	req := LineRequest{
		Points: []float32{100, 24, 5, 6},
		Tool: "tr",
	}
	expected := ";100_24;5_6"
	res := normalizeCoordinates(req)
	if res != expected {
		t.Fatalf("not match")
	}
}

func TestNormalizeCoordinateList(t *testing.T) {
	reqArray := make([]LineRequest, 0)
	req := LineRequest{
		Points: []float32{100, 24, 5, 6},
		Tool: "tr",
	}
	reqArray = append(reqArray, req)
	expected := "100_24;5_6"
	res := concatCoordinates(reqArray)
	if expected != res {
		t.Fatalf("not match")
	}
}

func TestNormalzeCoordinateWithTwoList(t *testing.T) {
	reqArray := make([]LineRequest, 0)
	req := LineRequest{
		Points: []float32{100, 24, 5, 6},
		Tool: "tr",
	}
	reqArray = append(reqArray, req)
	reqArray = append(reqArray, req) // append twice
	expected := "100_24;5_6;100_24;5_6"
	res := concatCoordinates(reqArray)
	if expected != res {
		t.Fatalf("not match")
	}
}