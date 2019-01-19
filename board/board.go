package board

type (
	Board struct {
		Height       int
		Width        int
		BallPosX     int
		BallPosY     int
		RacketAPos   int
		RacketBPos   int
		PlayerAScore int
		PlayerBScore int
	}
)

func NewBoard(boardSize int) *Board {
	return &Board{
		Height:       boardSize,
		Width:        boardSize * 3,
		PlayerAScore: 0,
		PlayerBScore: 0,
	}
}
