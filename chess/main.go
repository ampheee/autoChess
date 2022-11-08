package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/notnil/chess"
	"image/color"
	"log"
	"math/rand"
	"time"
)

func main() {
	a := app.New()
	mainPage := a.NewWindow("ChessGolang")
	game := chess.NewGame()
	grid := createGrid(game.Position().Board())
	over := canvas.NewImageFromResource(nil)
	over.FillMode = canvas.ImageFillContain
	over.Hide()
	mainPage.SetContent(container.NewMax(grid, container.NewWithoutLayout(over)))
	mainPage.Resize(fyne.NewSize(480, 480))
	mainPage.SetFixedSize(true)
	go func() {
		rand.Seed(time.Now().Unix())
		for game.Outcome() == chess.NoOutcome {
			time.Sleep(time.Millisecond * 500)
			moves := game.ValidMoves()
			mv := moves[rand.Intn(len(moves))]
			move(mv, game, grid, over)
		}
	}()
	mainPage.ShowAndRun()
}

func createGrid(board *chess.Board) *fyne.Container {
	var cells []fyne.CanvasObject
	//fill in background of grid. color settings
	for y := 7; y >= 0; y-- {
		for x := 0; x < 8; x++ {
			cell := canvas.NewRectangle(color.NRGBA{R: 0xF4, G: 0xE2, B: 0xB6, A: 0xFF})
			effect := canvas.NewImageFromResource(resourceWood1Png)
			if x%2 == y%2 {
				cell.FillColor = color.Gray{Y: 0x52}
				effect = canvas.NewImageFromResource(resourceWood2Png)
			}
			p := board.Piece(chess.Square(x + y*8))
			figure := canvas.NewImageFromResource(resourceForPiece(p))
			figure.FillMode = canvas.ImageFillContain
			cells = append(cells, container.NewMax(cell, effect, figure))
		}
	}
	return container.New(&boardLayout{}, cells...)
}

func refreshGrid(grid *fyne.Container, board *chess.Board) {
	y, x := 7, 0
	for _, cell := range grid.Objects {
		p := board.Piece(chess.Square(x + y*8))
		img := cell.(*fyne.Container).Objects[2].(*canvas.Image)
		img.Resource = resourceForPiece(p)
		img.Refresh()
		x += 1
		if x == 8 {
			x = 0
			y--
		}
	}
}

func move(mv *chess.Move, game *chess.Game, grid *fyne.Container, over *canvas.Image) {
	off := squareToOff(mv.S1())
	cell := grid.Objects[off].(*fyne.Container)
	img := cell.Objects[2].(*canvas.Image)
	pos1 := cell.Position()

	over.Resource = img.Resource
	over.Move(pos1)
	over.Resize(img.Size())

	over.Show()
	img.Resource = nil
	img.Refresh()

	off = squareToOff(mv.S2())
	cell = grid.Objects[off].(*fyne.Container)
	pos2 := cell.Position()

	a := canvas.NewPositionAnimation(pos1, pos2, time.Millisecond*500, func(p fyne.Position) {
		over.Move(p)
		over.Refresh()
	})
	a.Start()
	time.Sleep(time.Millisecond * 480)
	err := game.Move(mv)
	if err != nil {
		log.Printf("error! %s", err)
	}
	refreshGrid(grid, game.Position().Board())
	over.Hide()

}

func squareToOff(square chess.Square) int {
	x := square % 8
	y := 7 - ((square - x) / 8)
	return int(x + y*8)
}
