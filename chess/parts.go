//go:generate fyne bundle -o bundledParts.go assets/figures

package main

import (
	"fyne.io/fyne/v2"
	"github.com/notnil/chess"
)

func resourceForPiece(p chess.Piece) fyne.Resource {
	switch p.Color() {
	case chess.Black:
		switch p.Type() {
		case chess.Pawn:
			return resourceChessPawnBlackPng
		case chess.King:
			return resourceChessKingBlackPng
		case chess.Bishop:
			return resourceChessBishopBlackPng
		case chess.Knight:
			return resourceChessKnightBlackPng
		case chess.Rook:
			return resourceChessRookBlackPng
		case chess.Queen:
			return resourceChessQueenBlackPng
		}
	case chess.White:
		switch p.Type() {
		case chess.Pawn:
			return resourceChessPawnWhitePng
		case chess.King:
			return resourceChesskingWhitePng
		case chess.Bishop:
			return resourceChessBishopWhitePng
		case chess.Knight:
			return resourceChessKnightWhitePng
		case chess.Rook:
			return resourceChessRookWhitePng
		case chess.Queen:
			return resourceChessQueenWhitePng
		}
	}
	return nil
}
