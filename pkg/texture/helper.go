package texture

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/gps/pkg/pwd"
)

var WHITE = color.RGBA{255, 255, 255, 255}
var BLACK = color.RGBA{0, 0, 0, 255}

type Texture interface {
	Save(session string, folder string)
	Dimension() int
}

type texture struct {
	container
	currentX int   // the current pixel x to map color
	currentY int   // the current pixel y to map color
	code     string // the code of texture
}

type container struct {
	width  	int
	height 	int
	img  	*image.RGBA
	
}

func NewTexture(width int, height int) (Texture) {
	con := container{
		width: width,
		height: height,
		img: image.NewRGBA(image.Rect(0, 0, width, height)),
	}

	return &texture{
		container: con,
		currentX: 0,
		currentY: 0,
	}
}

/*
Return dimension of texture
*/
func (t *texture) Dimension() int {
	return t.container.width * t.container.height
}

/*
Generate texture code
*/
func (t *texture) generateCode() string {
	textLength := t.Dimension() / 8
	return pwd.GenerateText(textLength)
}

func (t *texture) generateCodeBinary(char rune) string {
	return pwd.TextToBinary(char)
}

func (t * texture) generatePixelColor(x int, y int, flag rune) {
	if flag == '0' {
		t.savePixel(x, y, WHITE)
	} else {
		t.savePixel(x, y, BLACK)
	}
}

func (t *texture) mutateColor(x int, y int) {

}


func (t *texture) savePixel(x int, y int, mapping color.RGBA) {
	t.container.img.Set(x, y, mapping)
}

func (t *texture) generateTexture() {
	code := t.generateCode()
	t.code = code
	for _, c := range code {
		binary := t.generateCodeBinary(c)
		for _, bin := range binary {
			t.generatePixelColor(t.currentX, t.currentY, bin)
			t.currentX += 1 // move x to next column
			if t.currentX == t.container.width {
				t.currentX = 0 // reset x value
				t.currentY += 1 // increase y by 1
			}
		}
	}
}

/*
Session name is required to name the image belongs to current session
A folder is required to put the generatd texture in
*/
func (t *texture) Save(session string, folder string) {
	t.generateTexture()
	dir := fmt.Sprintf("%s/%s", folder, session)
	os.Mkdir(dir, os.ModePerm)
	filename := fmt.Sprintf("%s/%s.png", dir, t.code)
	f, err := os.Create(filename)
	
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, t.container.img)
}