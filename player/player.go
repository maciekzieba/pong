package player

import (
	"github.com/maciekzieba/pong/object"
)

type Player struct {
	Name   string
	Racket *object.Racket
	Input  chan object.RacketVectorValue
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:   name,
		Racket: nil,
		Input:  make(chan object.RacketVectorValue),
	}
}
