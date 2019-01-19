package object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBallMove(t *testing.T) {
	ball := NewBall(5, 5, BallVectorLeftUp)
	ball.Move()
	assert.Equal(t, 4, ball.PosX)
	assert.Equal(t, 4, ball.PosY)

	ball = NewBall(5, 5, BallVectorUpRight)
	ball.Move()
	assert.Equal(t, 6, ball.PosX)
	assert.Equal(t, 4, ball.PosY)

	ball = NewBall(5, 5, BallVectorDownLeft)
	ball.Move()
	assert.Equal(t, 4, ball.PosX)
	assert.Equal(t, 6, ball.PosY)

	ball = NewBall(5, 5, BallVectorRightDown)
	ball.Move()
	assert.Equal(t, 6, ball.PosX)
	assert.Equal(t, 6, ball.PosY)
}

func TestBallHit(t *testing.T) {
	ball := NewBall(5, 5, BallVectorRightDown)
	ball.Hit(BallHitNone)
	assert.Equal(t, BallVectorRightDown, ball.Vector)

	ball = NewBall(5, 5, BallVectorRightDown)
	ball.Hit(BallHitDown)
	assert.Equal(t, BallVectorUpRight, ball.Vector)

	ball = NewBall(5, 5, BallVectorDownLeft)
	ball.Hit(BallHitDown)
	assert.Equal(t, BallVectorLeftUp, ball.Vector)

	ball = NewBall(5, 5, BallVectorLeftUp)
	ball.Hit(BallHitUp)
	assert.Equal(t, BallVectorDownLeft, ball.Vector)

	ball = NewBall(5, 5, BallVectorLeftUp)
	ball.Hit(BallHitUp)
	assert.Equal(t, BallVectorDownLeft, ball.Vector)

	ball = NewBall(5, 5, BallVectorLeftUp)
	ball.Hit(BallHitLeft)
	assert.Equal(t, BallVectorUpRight, ball.Vector)

	ball = NewBall(5, 5, BallVectorDownLeft)
	ball.Hit(BallHitLeft)
	assert.Equal(t, BallVectorRightDown, ball.Vector)

	ball = NewBall(5, 5, BallVectorRightDown)
	ball.Hit(BallHitRight)
	assert.Equal(t, BallVectorDownLeft, ball.Vector)

	ball = NewBall(5, 5, BallVectorUpRight)
	ball.Hit(BallHitRight)
	assert.Equal(t, BallVectorLeftUp, ball.Vector)
}
