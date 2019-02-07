package object

const (
	RacketVectorUp   racketVector = 0
	RacketVectorDown racketVector = 1
	RacketVectorNone racketVector = 2
	RacketSize                    = 5
	RacketPadding                 = 5
)

type (
	racketVector int

	RacketVectorValue interface {
		RacketVector() racketVector
	}

	Racket struct {
		Pos int
	}
)

func (rv racketVector) RacketVector() racketVector {
	return rv
}

func NewRacket(pos int) *Racket {
	return &Racket{
		Pos: pos,
	}
}
