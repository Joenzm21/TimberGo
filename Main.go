package main

import (
	"TimberGO/src"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	ebiten.SetMaxTPS(int(1 / src.DeltaTime))
	ebiten.SetWindowSize(int(src.ScreenWidth), int(src.ScreenHeight))
	ebiten.SetWindowTitle("Timber")
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(src.Init().New()); err != nil {
		log.Fatal(err)
	}
}
