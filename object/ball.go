package object

const (
	BallVectorUpRight   ballVector = 0
	BallVectorRightDown ballVector = 1
	BallVectorDownLeft  ballVector = 2
	BallVectorLeftUp    ballVector = 3

	BallHitNone  ballHit = 0
	BallHitLeft  ballHit = 1
	BallHitRight ballHit = 2
	BallHitUp    ballHit = 3
	BallHitDown  ballHit = 4
)

type (
	ballVector int
	ballHit    int

	BallHitValue interface {
		BallHit() ballHit
	}

	BallVectorValue interface {
		BallVector() ballVector
	}
	Ball struct {
		PosX   int
		PosY   int
		Vector ballVector
	}
)

func (bv ballVector) BallVector() ballVector {
	return bv
}

func (bh ballHit) BallHit() ballHit {
	return bh
}

func NewBall(posX int, posY int, vector ballVector) *Ball {
	return &Ball{
		PosX:   posX,
		PosY:   posY,
		Vector: vector,
	}
}

func (b *Ball) Hit(hitValue BallHitValue) {
	if hitValue == BallHitUp {
		if b.Vector == BallVectorLeftUp {
			b.Vector = BallVectorDownLeft
		} else if b.Vector == BallVectorUpRight {
			b.Vector = BallVectorRightDown
		}
	} else if hitValue == BallHitDown {
		if b.Vector == BallVectorDownLeft {
			b.Vector = BallVectorLeftUp
		} else if b.Vector == BallVectorRightDown {
			b.Vector = BallVectorUpRight
		}
	} else if hitValue == BallHitLeft {
		if b.Vector == BallVectorLeftUp {
			b.Vector = BallVectorUpRight
		} else if b.Vector == BallVectorDownLeft {
			b.Vector = BallVectorRightDown
		}
	} else if hitValue == BallHitRight {
		if b.Vector == BallVectorUpRight {
			b.Vector = BallVectorLeftUp
		} else if b.Vector == BallVectorRightDown {
			b.Vector = BallVectorDownLeft
		}
	}
}

func (b *Ball) Move() {
	switch b.Vector {
	case BallVectorLeftUp:
		b.PosX--
		b.PosY--
		break
	case BallVectorUpRight:
		b.PosX++
		b.PosY--
		break
	case BallVectorDownLeft:
		b.PosX--
		b.PosY++
		break
	case BallVectorRightDown:
		b.PosX++
		b.PosY++
		break
	}
}
