package auth

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gps/conf"
	"github.com/gps/pkg/texture"
	"github.com/gps/server/models"
)


const HEIGHT          = 16 // standard 16 pixels
const WIDTH           = 16
const FOLDER 		  = "./storage/sp" // storage folder to save texture
const NUM_TEXTURE     = 20 // number of texture to generate
const HOST            =  "http://localhost:8090" // hardcoded value for host url
const VERIFY_FOLDER   = "./storage/verify" // storage folder hold image for verify

type PWDTexture struct {
	URL  string `json:"url"` // image url
	ID   string `json:"id"`  // texture id
}

type Auth struct {
	session        string
	folder         string // folder to storage texture
	TextureWidth   int
	TextureHeight  int
	SecretKey      string
}

type LineRequest struct {
	Points []float32  `json:"points"`
	Tool   string `json:"tool"`
}

type RegisterRequest struct {
	Lines 	[]LineRequest `json:"lines"` // canvas x,y coordinate
	Order 	string  `json:"order"` // box order specification split by _ (underscore)
	Email   string  `json:"email"` // email of user
	Session string 	`json:"session"` // the session of this request
}

func NewAuth() (*Auth) {
	return &Auth{
		session: NewSession(),
		folder: FOLDER, // set as default
		TextureWidth: WIDTH,
		TextureHeight: HEIGHT,
		SecretKey: conf.EnvironmentSetting.Key,
	}
}

/*
Set texture folder destination, override default folder
*/
func (a *Auth) SetFolder(pathname string) {
	a.folder = pathname
}

func (a *Auth) RegisterPWD() []PWDTexture {
	return a.generateTextures(NUM_TEXTURE)
}

func (a *Auth) Session() string {
	return a.session
}

func (a *Auth) shuffleTextures(textures []PWDTexture) []PWDTexture{
	textureLength := len(textures)
	rand.Seed(time.Now().UnixNano())

	for i := textureLength -1; i > 0; i-- {
             
		// Pick a random index from 0 to i
		j := rand.Intn(i + 1)
		rand.Seed(time.Now().UnixNano())
	
		// Swap i with the element at random index
		temp := textures[i];
		textures[i] = textures[j];
		textures[j] = temp;
	}
	return textures
}

/*
Generate texture given number of texture to generate
*/
func (a *Auth) generateTextures(n int) []PWDTexture {
	textures := make([]PWDTexture, n)
	for i := 0; i < n; i++ {
		tex := texture.NewTexture(HEIGHT, WIDTH)
		tex.SetKey(a.SecretKey)
		tex.Save(a.session, FOLDER) // save texture temporariy to storage

		textureResponse := PWDTexture {
			URL: fmt.Sprintf("%s/cdn/%s/%s", HOST, a.session, tex.Code()),
			ID: tex.ID(),
		}
		textures[i] = textureResponse
	}
	textures = a.shuffleTextures(textures)

	return textures
	
}


/*
Possibly future implementation
After certain period, session will be invalidated, images are then removed from
hard disk. 
*/
func (a *Auth) ClearSessionTexture(sessionId string) {
	folder := fmt.Sprintf("%s/%s", FOLDER, sessionId)
	err := os.RemoveAll(folder)
	if err != nil {
		panic(err)
	}
	/*files, err := filepath.Glob(folder)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}*/
}

func roundup(n float32) int{
	return int( n +0.5) 
}

/*
x and y coordinate are concat into x_y format
*/
func normalizeCoordinates(line LineRequest) string {
	str := ""
	for idx, point := range line.Points {
		p := roundup(point)
		if idx % 2 == 0 { // x coordinate
			str += fmt.Sprintf(";%d", p) // add semicolon to the start
		} else if idx % 2 == 1 { // y coordinate
			str += fmt.Sprintf("_%d", p) // split point by underscore
		}
		/*
		if idx == 0 { // first inital point
			str += fmt.Sprintf("%d", p)
		} else {
			str += fmt.Sprintf("_%d", p) // split point by underscore
		}*/
	}
	return str
}

/*
Continuous Coordinates are concat to a string split by semicolon
*/
func concatCoordinates(lines []LineRequest) string {
	str := ""
	for idx, line := range lines {
		points := normalizeCoordinates(line)
		if idx == 0 { // first initial coordinate strip out 0
			points = points[1:]
		}
		str += points
		/*else {
			str += fmt.Sprintf("%s;", points) // split by semicolon
		}*/
	}
	return str
}

/*
Remove underscore, concat them all
*/
func concatOrder(orders string) string {
	orderArray := strings.Split(orders, "_")
	return strings.Join(orderArray[:], "")
}

/*
Register the user to db

The function will hash order and do all encryption methods
@return error
*/
func (a *Auth) Register(request RegisterRequest) error {
	coordinates := concatCoordinates(request.Lines)
	orders := concatOrder(request.Order)
	// hash order to store inside db
	
	hashedOrders := Hash(orders)
	encryptedCoordinate := EncryptCoordinates(coordinates, orders)

	insertErr := CreateUser(request.Email, hashedOrders, encryptedCoordinate)
	return insertErr
}

/*
This function must not tell whether the user password is correct

*/
func (a *Auth) VerifyUser(email string, orders string, session string) {
	var user models.User
	orders = concatOrder(orders)
	conf.DB.Where("email = ?", email).First(&user) // find by email
	if HasHash(user.Password, orders) { // check if password match
		text := DecryptCoordinates(user.Casting, orders) // decrypt coordinates
		photo := texture.NewPhoto(text)
		photo.GeneratePhoto()
		photo.Save(session, VERIFY_FOLDER, "1")
	}
}

/*
Except email
Box order and coordinate must not be in plaintext
Add user to database!
*/
func CreateUser(email string, hashedOrder string, encryptedCoordinates string) error {
	user := models.User{Email: email, Password: hashedOrder, Casting: encryptedCoordinates}
	err := user.Create()
	if err != nil {
		return err
	}
	return nil
}