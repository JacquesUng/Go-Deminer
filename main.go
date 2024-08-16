package main

import (
	"jung/deminer/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	height := 20
	width := 20
	game := game.NewGame(height, width)
	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}