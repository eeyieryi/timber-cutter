package main

import (
	"log"

	"github.com/eeyieryi/timber-cutter/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle(game.Title)
	if err := ebiten.RunGame(&game.Game{}); err != nil {
		if err == game.ErrTerminated {
			return
		}
		log.Fatal(err)
	}
}
