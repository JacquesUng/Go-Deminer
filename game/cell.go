package game

import "github.com/hajimehoshi/ebiten/v2"

type Cell struct {
	x           int
	y           int
	mine        bool
	exposed     bool
	minesAround int
	sprites     []*ebiten.Image
}