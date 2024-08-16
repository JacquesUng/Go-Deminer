package assets

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed *
var assets embed.FS

var OneSprite = mustLoadImage("one.png")
var TwoSprite = mustLoadImage("two.png")
var ThreeSprite = mustLoadImage("three.png")
var FourSprite = mustLoadImage("four.png")
var FiveSprite = mustLoadImage("five.png")
var SixSprite = mustLoadImage("six.png")
var SevenSprite = mustLoadImage("seven.png")
var EightSprite = mustLoadImage("eight.png")
var MineSprite = mustLoadImage("mine.png")
var UnopenedCellSprite = mustLoadImage("unopened-cell.png")
var OpenedCellSprite = mustLoadImage("opened-cell.png")

var Font = mustLoadFont("font.ttf")

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func mustLoadFont(name string) font.Face {
	f, err := assets.ReadFile(name)
	if err != nil {
		panic(err)
	}

	tt, err := opentype.Parse(f)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{Size: 48, DPI: 72, Hinting: font.HintingVertical})
	if err != nil {
		panic(err)
	}

	return face
}