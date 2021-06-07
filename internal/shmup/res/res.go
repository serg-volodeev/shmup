package res

import (
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Res struct {
	images map[string]*ebiten.Image
}

func LoadRes() *Res {
	r := &Res{}
	r.images = make(map[string]*ebiten.Image)
	r.images["bullet"] = loadImageFromFile("./assets/images/laserRed16.png")
	r.images["ship"] = loadImageFromFile("./assets/images/playerShip1_orange.png")
	r.images["space"] = loadImageFromFile("./assets/images/starfield.png")
	r.images["meteor1"] = loadImageFromFile("./assets/images/meteorBrown_big1.png")
	r.images["meteor2"] = loadImageFromFile("./assets/images/meteorBrown_med1.png")
	r.images["meteor3"] = loadImageFromFile("./assets/images/meteorBrown_small1.png")
	r.images["meteor4"] = loadImageFromFile("./assets/images/meteorBrown_tiny1.png")
	return r
}

func (r *Res) GetImage(name string) *ebiten.Image {
	return r.images[name]
}

func loadImageFromFile(path string) *ebiten.Image {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	img2 := ebiten.NewImageFromImage(img)
	return img2
}
