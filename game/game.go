package game

import (
	"fmt"
	"image/color"
	"jung/deminer/assets"
	"jung/deminer/utils"
	"math/rand"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Game struct {
	cells    [][]Cell
	height   int
	width    int
	exposed  	 int
	nMines int
	gameOver bool
}

func NewGame(height, width int) *Game {
	cells := make([][]Cell, 0)
	for y:=0 ; y<height; y++ {
		row := make([]Cell, 0)
		for x:=0; x<width; x++ {
			cellSprites := []*ebiten.Image{ assets.UnopenedCellSprite }
			cell := Cell{
				x: x,
				y: y,
				sprites: cellSprites,
			}
			row = append(row, cell)
		}
		cells = append(cells, row)
	}
	return &Game{
		cells: cells,
		height: height,
		width: width,
	}
}

func (game *Game) Update() error {
	if !game.gameOver {
		if game.exposed == game.height * game.width - game.nMines {
			fmt.Println("You won !")
			fmt.Println("Press space to restart.")
			game.gameOver = true
		} else {
			cellHeight := float64(utils.ScreenHeight) / float64(game.height)
			cellWidth := float64(utils.ScreenWidth) / float64(game.width)
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				xPx, yPx := ebiten.CursorPosition()
				x := xPx / int(cellWidth)
				y := yPx / int(cellHeight)
				if game.exposed == 0 {
					game.cells[y][x].mine = false
					game.nMines = game.generateMines(x, y)
				}
				game.handleCell(x, y)
			}
		}
	} else {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			fmt.Println("Restarting")
			game.Reset()
		}
	}
	return nil
}

func (game *Game) generateMines(x, y int) int {
	fmt.Println("Adding mines after clicking on " + fmt.Sprint(x) + ", " + fmt.Sprint(y))
	originalCellIndex := y*game.width+x
	nMines := game.height * game.width / 10
	fmt.Println("Generating " + fmt.Sprint(nMines) + " mines")
	cellsAroundIndexList := game.cellsAroundIndexList(originalCellIndex)
	fmt.Println("Index must not be part of")
	fmt.Println(cellsAroundIndexList)
	fmt.Println("Nor equal to")
	fmt.Println(originalCellIndex)
	mineIndexList := make([]int, 0)
	for len(mineIndexList) < nMines {
		randomIndex := rand.Intn(game.height*game.width)
		if randomIndex != originalCellIndex && !slices.Contains(cellsAroundIndexList, randomIndex) && !slices.Contains(mineIndexList, randomIndex) {
			mineIndexList = append(mineIndexList, randomIndex)
			mineY := randomIndex / game.width
			mineX := randomIndex % game.width
			game.cells[mineY][mineX].mine = true
			fmt.Println("Adding index: " + fmt.Sprint(randomIndex))	
			// fmt.Println("Adding mine : " + fmt.Sprint(mineX) + ", " + fmt.Sprint(mineY))
		}
	}
	return nMines
}

func (game *Game) handleCell(x, y int) {
	if x >= game.width || x < 0 || y >= game.height || y < 0 || game.cells[y][x].exposed {
		return
	}
	game.cells[y][x].exposed = true
	game.exposed++
	if game.cells[y][x].mine {
		fmt.Println("Stepped on a mine, GAME OVER.")
		fmt.Println("Click on esc to restart")
		cellSprite := assets.OpenedCellSprite
		mineSprite := assets.MineSprite
		game.cells[y][x].sprites[0] = cellSprite
		game.cells[y][x].sprites = append(game.cells[y][x].sprites, mineSprite)
		game.gameOver = true
		return
	}
	cellsAround := game.cellsAround(x, y)
	minesAround := 0
	for _, cell := range cellsAround {
		if cell.mine {
			minesAround++
		}
	}
	game.cells[y][x].minesAround = minesAround

	if minesAround == 0 {
		game.cells[y][x].sprites[0] = assets.OpenedCellSprite
		for _, cell := range cellsAround {
			if !cell.exposed {
				game.handleCell(cell.x, cell.y)
			}
		}
	} else {
		switch minesAround {
		case 1:
			game.cells[y][x].sprites[0] = assets.OneSprite
		case 2:
			game.cells[y][x].sprites[0] = assets.TwoSprite
		case 3:
			game.cells[y][x].sprites[0] = assets.ThreeSprite
		case 4:
			game.cells[y][x].sprites[0] = assets.FourSprite
		case 5:
			game.cells[y][x].sprites[0] = assets.FiveSprite
		case 6:
			game.cells[y][x].sprites[0] = assets.SixSprite
		case 7:
			game.cells[y][x].sprites[0] = assets.SevenSprite
		case 8:
			game.cells[y][x].sprites[0] = assets.EightSprite
		}
	}
}

func (game *Game) cellsAroundIndexList(cellIndex int) []int {
	return []int{
		cellIndex-game.width-1,
		cellIndex-game.width,
		cellIndex-game.width+1,
		cellIndex-1,
		cellIndex+1,
		cellIndex+game.width-1,
		cellIndex+game.width,
		cellIndex+game.width+1,
	}
}

func (game *Game) cellsAround(x, y int) []*Cell {
	cellsAround := make([]*Cell, 0)
	if x > 0 && y > 0 {
		cellsAround = append(cellsAround, &game.cells[y-1][x-1])
	}
	if y > 0 {
		cellsAround = append(cellsAround, &game.cells[y-1][x])
	}
	if x < game.width-1 && y > 0 {
		cellsAround = append(cellsAround, &game.cells[y-1][x+1])
	}
	if x > 0 {
		cellsAround = append(cellsAround, &game.cells[y][x-1])
	}
	if x < game.width-1 {
		cellsAround = append(cellsAround, &game.cells[y][x+1])
	}
	if x > 0 && y < game.height-1 {
		cellsAround = append(cellsAround, &game.cells[y+1][x-1])
	}
	if y < game.height-1 {
		cellsAround = append(cellsAround, &game.cells[y+1][x])
	}
	if x < game.width-1 && y < game.height-1 {
		cellsAround = append(cellsAround, &game.cells[y+1][x+1])
	}
	return cellsAround
}

func (game *Game) Draw(screen *ebiten.Image) {
	cellHeight := float64(utils.ScreenHeight) / float64(game.height)
	cellWidth := float64(utils.ScreenWidth) / float64(game.width)

	for x:=0; x<game.width; x++ {
		for y:=0; y<game.height; y++ {
			cellSprites := game.cells[y][x].sprites
			cellBounds := cellSprites[0].Bounds()
			op := &ebiten.DrawImageOptions{}

			op.GeoM.Scale(cellWidth / float64(cellBounds.Dx()), cellHeight / float64(cellBounds.Dy()))
			op.GeoM.Translate(float64(x)*cellWidth, float64(y)*cellHeight)
			screen.DrawImage(cellSprites[0], op)

			if len(cellSprites) == 2 {
				secondaryBounds := cellSprites[1].Bounds()
				opSecondary := &ebiten.DrawImageOptions{}
				opSecondary.GeoM.Scale(cellWidth / float64(secondaryBounds.Dx()), cellHeight / float64(secondaryBounds.Dy()))
				opSecondary.GeoM.Translate(float64(x)*cellWidth, float64(y)*cellHeight)
				screen.DrawImage(cellSprites[1], opSecondary)
			}
		}
	}

	textToPrint := ""
	if !game.gameOver {
		textToPrint = "Cells exposed: " + fmt.Sprint(game.exposed) + "/" + fmt.Sprint(game.height*game.width-game.nMines)
	} else {
		textToPrint = "Press space to restart"
	}
	
	textX := utils.ScreenWidth / 10
	textY := utils.ScreenHeight + utils.ScreenHeight / 10
	text.Draw(screen, textToPrint, assets.Font, textX, textY, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return utils.ScreenWidth, utils.ScreenHeight+100
}

func (game *Game) Reset() {
	newGame := NewGame(game.height, game.width)
	game.cells = newGame.cells
	game.gameOver = false
	game.exposed = 0
}