package shmup

import (
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Res struct {
	images map[string]*ebiten.Image
}

func loadRes() *Res {
	r := &Res{}
	r.images = make(map[string]*ebiten.Image)
	r.images["bullet"] = loadImageFromFile("./assets/images/laserRed16.png")
	r.images["ship"] = loadImageFromFile("./assets/images/playerShip1_orange.png")
	r.images["space"] = loadImageFromFile("./assets/images/starfield.png")
	return r
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
