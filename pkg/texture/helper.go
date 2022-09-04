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
	SetKey(key string)
	ID() string
	SetCode(code string)
	Code() string
}

type texture struct {
	container
	currentX int    // the current pixel x to map color
	currentY int    // the current pixel y to map color
	code     string // the code of texture
	key      string // secret key
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

/*
Set pixel color at x, y coordinates
If flag is 0, pixel will be white
flag is 1, pixel will be black
*/
func (t * texture) generatePixelColor(x int, y int, flag rune) {
	if flag == '0' {
		t.savePixel(x, y, WHITE)
	} else {
		t.savePixel(x, y, BLACK)
	}
}

func (t *texture) mutateColor(x int, y int) {

}

/*
Override texture code, unless necessary, do not use
*/
func (t *texture) SetCode(code string) {
	t.code = code
}

/*
Set app secret key
*/
func (t *texture) SetKey(key string) {
	t.key = key
}

func sliceKey(key string) string {
	return key[0:7]
}

func textToBinary(text string) string {
	bin := ""
	for _, c := range text {
		binary := pwd.TextToBinary(c)
		bin += binary
	}

	return bin
}

func XOR(a rune, b rune) rune {
	if a == b {
		return '0'
	}
	return '1'
}

/*
Return ID of the texture, id is different from texture code
*/
func (t *texture) ID() string {
	if len(t.key) < 7 {
		panic("Key not long enough")
	}

	if len(t.code) < 7 {
		panic("Code not initialized yet")
	}
	key := sliceKey(t.key)
	keyBin := textToBinary(key)

	textureCode := sliceKey(t.code) // slice to first 7 segment
	textureBin := textToBinary(textureCode)

	xorPattern := ""
	for index, i := range keyBin {
		xorResult := XOR(i, rune(textureBin[index]))
		xorPattern += string(xorResult)
	}
	
	id, err := pwd.BinaryToHex(xorPattern)
	if err != nil {
		panic(err)
	}

	return id
	
}

func (t *texture) Code() string {
	return t.code
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