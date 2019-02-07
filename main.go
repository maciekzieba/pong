package main

import (
	"github.com/maciekzieba/pong/game"
	"github.com/maciekzieba/pong/io"
	"github.com/maciekzieba/pong/player"
	"github.com/maciekzieba/pong/ui/asci"
)

func main() {
	pA := player.NewPlayer("A")
	pB := player.NewPlayer("B")

	keyboard := io.NewKeyboard(pA.Input, pB.Input)
	go keyboard.ListenToKeyboard()

	asciDrawer := asci.NewAsci()
	match := game.NewMatch(pA, pB)
	match.StartMatch(asciDrawer)
}
