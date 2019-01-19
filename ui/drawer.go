package ui

import "github.com/maciekzieba/pong/board"

type Drawer interface {
	Draw(boardToDraw *board.Board)
}
