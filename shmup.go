package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/serg-volodeev/shmup/internal/shmup"
)

func main() {
	game := shmup.NewGame()
	ebiten.SetWindowSize(shmup.ScreenWidth, shmup.ScreenHeight)
	ebiten.SetWindowTitle(shmup.Title)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
