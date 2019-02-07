package io

import (
	"github.com/maciekzieba/pong/object"
	"github.com/nsf/termbox-go"
	"os"
)

type Keyboard struct {
	playerAChannel chan object.RacketVectorValue
	playerBChannel chan object.RacketVectorValue
}

func NewKeyboard(pA chan object.RacketVectorValue, pB chan object.RacketVectorValue) *Keyboard {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	return &Keyboard{
		playerAChannel: pA,
		playerBChannel: pB,
	}
}

func (k *Keyboard) ListenToKeyboard() {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowDown:
				k.playerAChannel <- object.RacketVectorDown
			case termbox.KeyArrowUp:
				k.playerAChannel <- object.RacketVectorUp
			case termbox.KeyEsc:
				termbox.Close()
				os.Exit(0)
			}
			switch ev.Ch {
			case 119:
				k.playerBChannel <- object.RacketVectorUp
			case 115:
				k.playerBChannel <- object.RacketVectorDown
			}
		}
	}
}
