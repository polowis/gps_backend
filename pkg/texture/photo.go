package texture

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
)

const CANVAS_WIDTH  = 300
const CANVAS_HEIGHT = 300

type Photo struct {
	img         *image.RGBA
	coordinates string
}

// Pixel
type Pixel struct {
    R int
    G int
    B int
    A int
}

/*
The provided coordinate must be a string representation
x1_y1;x2_y2
*/
func NewPhoto(coordinates string) (*Photo) {
	con := Photo{
		img: image.NewRGBA(image.Rect(0, 0, CANVAS_WIDTH, CANVAS_HEIGHT)),
		coordinates: coordinates,
	}
	return &con
}

func (p *Photo) Width() int {
	return CANVAS_WIDTH
}

func (p *Photo) Height() int {
	return CANVAS_HEIGHT
}

/*
Return dimension of photo
*/
func (p *Photo) Dimension() int {
	return CANVAS_WIDTH * CANVAS_HEIGHT
}

func (p *Photo) savePixel(x int, y int, mapping color.RGBA) {
	p.img.Set(x, y, mapping)
}

func (p *Photo) splitCoordinates() []string {
	coordinatesList := strings.Split(p.coordinates, ";") // split by semicolon
	return coordinatesList
}

func (p *Photo) splitXYCoordinates(coordinates string) (int, int) {
	coordinatesList := strings.Split(coordinates, "_") // split x and y

	xString := coordinatesList[0]
	yString := coordinatesList[1] // hardcoded index

	x, _ := strconv.Atoi(xString)
	y, _ := strconv.Atoi(yString)
	return x, y

}

func (p *Photo) GeneratePhoto() {
	groupCoordinates := p.splitCoordinates()
	for _, points := range groupCoordinates {
		x, y := p.splitXYCoordinates(points)
		p.savePixel(x, y, BLACK)
	}
	p.makeWhite()
}

/*
Any pixel that transparent (not init yet) will have to be white
*/
func (p *Photo) makeWhite() {
	for y := 0; y < CANVAS_HEIGHT; y++ {
        for x := 0; x < CANVAS_WIDTH; x++ {
            pixel := rgbaToPixel(p.img.At(x, y).RGBA())
			if pixel.A == 0 { // if opacity is 0
				p.savePixel(x, y, WHITE)
			}
        }
    }
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
    return Pixel{
		int(r / 257), 
		int(g / 257), 
		int(b / 257), 
		int(a / 257),
	}
}

/*
Session name is required to name the image belongs to current session
A folder is required to put the generatd texture in
*/
func (p *Photo) Save(session string, folder string, id string) {
	dir := fmt.Sprintf("%s/%s", folder, session)
	os.Mkdir(dir, os.ModePerm)
	filename := fmt.Sprintf("%s/%s.png", dir, id)
	f, err := os.Create(filename)
	
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, p.img)
}