package object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBallServe(t *testing.T) {
	ball := NewBall(5, 5, BallVectorLeftUp)
	ball.Serve(ServeSideLeft, 1, 1)
	assert.Equal(t, 1, ball.PosX)
	assert.Equal(t, 1, ball.PosY)
	assert.Equal(t, BallVectorRightDown, ball.Vector)

	ball = NewBall(5, 5, BallVectorLeftUp)
	ball.Serve(ServeSideRight, 10, 10)
	assert.Equal(t, 10, ball.PosX)
	assert.Equal(t, 10, ball.PosY)
	assert.Equal(t, BallVectorDownLeft, ball.Vector)
}

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
