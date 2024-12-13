package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/uncleBlobby/infinigong/internal/pkg/gong"
)

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("infinigong")
	g := gong.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
