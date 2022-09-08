package texture

import "testing"

func TestSimplePhotoCoordinates(t *testing.T) {
	coordinates := "12_34;56_12;90_13"
	p := NewPhoto(coordinates)
	p.GeneratePhoto()
	p.Save("1234", "../../storage/verify", "1")
}