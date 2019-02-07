package player

import (
	"github.com/maciekzieba/pong/io"
	"github.com/maciekzieba/pong/object"
)

type Player struct {
	Name       string
	Racket     *object.Racket
	Controller io.Controller
}

func NewPlayer(name string, controller io.Controller) *Player {
	return &Player{
		Name:       name,
		Racket:     nil,
		Controller: controller,
	}
}
