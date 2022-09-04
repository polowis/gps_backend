package auth

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gps/conf"
	"github.com/gps/pkg/texture"
)


const HEIGHT          = 16 // standard 16 pixels
const WIDTH           = 16
const FOLDER 		  = "./storage/sp" // storage folder to save texture
const NUM_TEXTURE     = 20 // number of texture to generate
const HOST            =  "http://localhost:8090" // hardcoded value for host url

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