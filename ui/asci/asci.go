package asci

import (
	"fmt"

	tm "github.com/buger/goterm"
	"github.com/maciekzieba/pong/board"
	"github.com/maciekzieba/pong/object"
)

type Asci struct {
}

func NewAsci() *Asci {
	return &Asci{}
}

func (a *Asci) Draw(b *board.Board) {
	tm.Clear()
	tm.MoveCursor(1, 1)
	//Board
	box := tm.NewBox(b.Width, b.Height, 0)
	tm.Print(box)

	//Ball
	tm.Print(tm.MoveTo("*", b.BallPosX, b.BallPosY))

	//Rackets
	for i := 0; i < object.RacketSize; i++ {
		tm.Print(tm.MoveTo("#", object.RacketPadding, b.RacketAPos+i))
		tm.Print(tm.MoveTo("#", b.Width-object.RacketPadding, b.RacketBPos+i))
	}

	//Points
	tm.MoveCursor(1, b.Height+1)
	tm.Print(tm.Color(fmt.Sprintf("%d : %d", b.PlayerAScore, b.PlayerBScore), tm.RED))

	tm.MoveCursor(1, b.Height+2)
	tm.Print(fmt.Sprintf("Ball x: %d y:%d", b.BallPosX, b.BallPosY))
	tm.MoveCursor(1, b.Height+3)
	tm.Print(fmt.Sprintf("PlayerA:%d PlayerB:%d", b.RacketAPos, b.RacketBPos))
	tm.Flush()

}
