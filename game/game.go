package game

import (
	"time"

	"github.com/maciekzieba/pong/board"
	"github.com/maciekzieba/pong/object"
	"github.com/maciekzieba/pong/player"
	"github.com/maciekzieba/pong/ui"
)

const (
	Waiting  GameStatus = 0
	OnGoing  GameStatus = 1
	Finished GameStatus = 2

	BoardSizeSmall  = 15
	BoardSizeMedium = 20
	BoardSizeLarge  = 25
)

type (
	Match struct {
		ID        int
		PlayerA   *player.Player
		PlayerB   *player.Player
		BoardSize int
		Status    GameStatus
		Drawer    ui.Drawer
		Board     *board.Board
		Ball      *object.Ball
	}

	GameStatus int
)

func NewMatch(playerA *player.Player, playerB *player.Player) *Match {
	playerA.Racket = object.NewRacket(2)
	playerB.Racket = object.NewRacket(2)

	match := &Match{
		PlayerA:   playerA,
		PlayerB:   playerB,
		Status:    Waiting,
		BoardSize: BoardSizeMedium,
	}
	return match
}

func (m *Match) StartMatch(drawer ui.Drawer) {
	m.Drawer = drawer
	m.Status = OnGoing
	m.Ball = object.NewBall(10, 5, object.BallVectorRightDown)
	m.Board = board.NewBoard(m.BoardSize)
	m.loop()
}

func (m *Match) MovePlayerRacket(player *player.Player, vector object.RacketVectorValue) {
	if vector == object.RacketVectorUp {
		player.Racket.Pos--
	} else if vector == object.RacketVectorDown {
		player.Racket.Pos++
	}
}

func (m *Match) loop() {
	tick := time.Tick(50 * time.Millisecond)
	ballTick := time.Tick(250 * time.Millisecond)
	for {
		select {
		case <-tick:
			m.refreshBoard()
			break
		case <-ballTick:
			m.refreshBall()
		}
	}
}

func (m *Match) detectBandCollision() object.BallHitValue {
	if m.Board.BallPosY <= 1 {
		return object.BallHitUp
	}
	if m.Board.BallPosY >= m.Board.Height {
		return object.BallHitDown
	}
	return object.BallHitNone
}

func (m *Match) detectRacketCollision() object.BallHitValue {
	if m.Ball.Vector == object.BallVectorDownLeft || m.Ball.Vector == object.BallVectorLeftUp {
		racket := m.PlayerA.Racket
		racketPosX := object.RacketPadding
		if racketPosX+1 == m.Ball.PosX && (m.Ball.PosY >= racket.Pos && m.Ball.PosY <= (racket.Pos+object.RacketSize)) {
			return object.BallHitLeft
		}
	} else {
		racket := m.PlayerB.Racket
		racketPosX := m.Board.Width - object.RacketPadding
		if racketPosX-1 == m.Ball.PosX && (m.Ball.PosY >= racket.Pos && m.Ball.PosY <= (racket.Pos+object.RacketSize)) {
			return object.BallHitRight
		}
	}

	return object.BallHitNone
}

func (m *Match) detectPointLoss() {
	if m.Ball.PosX <= 0 {
		m.Board.PlayerBScore++
		servePosX := m.Board.Width - object.RacketPadding - 1
		servePosY := m.PlayerB.Racket.Pos
		m.Ball.Serve(object.ServeSideRight, servePosX, servePosY)
	} else if m.Ball.PosX >= m.Board.Width {
		m.Board.PlayerAScore++
		servePosX := object.RacketPadding + 1
		servePosY := m.PlayerA.Racket.Pos
		m.Ball.Serve(object.ServeSideLeft, servePosX, servePosY)
	}
}

func (m *Match) refreshBall() {
	m.Ball.Move()
}

func (m *Match) refreshBoard() {
	m.Board.BallPosX = m.Ball.PosX
	m.Board.BallPosY = m.Ball.PosY
	m.Board.RacketAPos = m.PlayerA.Racket.Pos
	m.Board.RacketBPos = m.PlayerB.Racket.Pos

	m.Ball.Hit(m.detectBandCollision())
	m.Ball.Hit(m.detectRacketCollision())
	m.detectPointLoss()

	m.Drawer.Draw(m.Board)
}
