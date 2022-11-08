//go:generate fyne bundle -o bundledBoard.go assets/cells

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type boardLayout struct{}

func (b *boardLayout) Layout(cells []fyne.CanvasObject, size fyne.Size) {
	smallEdge := size.Width
	if size.Height < size.Width {
		smallEdge = size.Height
	}

	leftInsert := (size.Width - smallEdge) / 2
	cellEdge := smallEdge / 8
	cellSize := fyne.NewSize(cellEdge, cellEdge)
	i := 0
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			cells[i].Resize(cellSize)
			cells[i].Move(fyne.NewPos(leftInsert+(float32(x)*cellEdge), float32(y)*cellEdge))
			i++
		}
	}
}

func (b *boardLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	edge := theme.IconInlineSize()
	return fyne.NewSize(edge, edge)
}
