package main

import (
	"fmt"

	"github.com/maciekzieba/pong/game"
	"github.com/maciekzieba/pong/io/keyboard"
	"github.com/maciekzieba/pong/matcher"
	"github.com/maciekzieba/pong/player"
	"github.com/maciekzieba/pong/ui/asci"
)

func main() {
	controllerA := keyboard.NewKeyboardController()
	controllerB := keyboard.NewKeyboardController()
	pA := player.NewPlayer("A", controllerA)
	pB := player.NewPlayer("B", controllerB)

	pMatcher := matcher.NewPlayerMatcher()
	pMatcher.AddPlayer(pA)
	pMatcher.AddPlayer(pB)

	pair := pMatcher.TryMatch()
	fmt.Printf("%s %s", pair.PlayerA.Name, pair.PlayerB.Name)

	asciDrawer := asci.NewAsci()
	match := game.NewMatch(pair.PlayerA, pair.PlayerB)
	match.StartMatch(asciDrawer)
}
