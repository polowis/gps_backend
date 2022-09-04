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
