package texture

import (
	"fmt"
	"testing"
)

func TestLinearEquationInterceptY(t *testing.T) {
	src := Point {
		x: -5,
		y: 10,
	}

	dest := Point {
		x: -3,
		y: 4,
	}

	slope, interceptY := useLinearEquation(src, dest)
	if slope != -3 {
		t.Fatalf("Slope not correct")
	}

	if interceptY != -5 {
		fmt.Println(interceptY)
		t.Fatalf("intercept not correct")
	}
}