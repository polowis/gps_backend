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
	points      []Point 
}

type Point struct {
	x int
	y int
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
		points: make([]Point, 0),
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

func (p *Photo) saveMarkedPoint(point Point) {
	p.points = append(p.points, point)
}

func (p *Photo) GeneratePhoto() {
	groupCoordinates := p.splitCoordinates()
	for _, points := range groupCoordinates {
		x, y := p.splitXYCoordinates(points)
		x = x % 300 // scale image
		y = y % 300

		// create point struct
		point := Point{
			x: x,
			y: y,
		}
		p.savePixel(x, y, BLACK)
		p.saveMarkedPoint(point)
	}
	p.makeWhite()
}

func squareDistance(x1 int, y1 int, x2 int, y2 int) int {
	xDiff := x1 - x2
	yDiff := y1 - y2
	return (xDiff * xDiff + yDiff * yDiff)

}

func (p *Photo) findNearestPoint(x int, y int) Point {
	closestPoint := p.points[0] // assume first point is closest
	shortestPath := squareDistance(x, y, p.points[0].x, p.points[1].y)
	for _, point := range p.points {
		// if same coordinate skip it, we do not want to map it to its own value
		if point.x == x && point.y == y {
			continue
		}

		distance := squareDistance(x, y, point.x, point.y)
		if distance < shortestPath { // if distance is smaller than flagged distance
			closestPoint = point  // set as new shortest point
			shortestPath = distance
		}
	}
	return closestPoint
}

func (p *Photo) LinearDraw(srcPoint Point, targetPoint Point) {
	if srcPoint.x < targetPoint.x {
		for i := srcPoint.x; i < targetPoint.x; i++ {
			p.savePixel(i, srcPoint.y, BLACK)
		}
	} else {
		for i := targetPoint.x; i < srcPoint.x; i++ {
			p.savePixel(i, srcPoint.y, BLACK)
		}
	}
	
	if srcPoint.y < targetPoint.y {
		for i := srcPoint.y; i < targetPoint.y; i++ {
			p.savePixel(targetPoint.x, i, BLACK)
		}
	} else {
		for i := targetPoint.y; i < srcPoint.y; i++ {
			p.savePixel(targetPoint.x, i, BLACK)
		}
	}
}

func (p *Photo) Smoothen() {
	for y := 0; y < CANVAS_HEIGHT; y++ {
        for x := 0; x < CANVAS_WIDTH; x++ {
            pixel := rgbaToPixel(p.img.At(x, y).RGBA())
			srcPoint := Point {
				x: x,
				y: y,
			}
			if isMarkedWith(pixel, BLACK) {
				targetPoint := p.findNearestPoint(x, y)
				p.LinearDraw(srcPoint, targetPoint)
				
			}
        }
    }
}

/*
Return true if the x y pixel is marked with color
*/
func isMarkedWith(pixel Pixel, color color.RGBA) bool {
	return pixel.R == int(color.R) && pixel.B == int(color.B) && pixel.G == int(color.G) && pixel.A == int(color.A)
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