package main

import (
	"fmt"

	"github.com/maciekzieba/pong/game"
	"github.com/maciekzieba/pong/matcher"
	"github.com/maciekzieba/pong/player"
	"github.com/maciekzieba/pong/ui/asci"
)

func main() {
	pA := player.NewPlayer("A")
	pB := player.NewPlayer("B")

	pMatcher := matcher.NewPlayerMatcher()
	pMatcher.AddPlayer(pA)
	pMatcher.AddPlayer(pB)

	pair := pMatcher.TryMatch()
	fmt.Printf("%s %s", pair.PlayerA.Name, pair.PlayerB.Name)

	asciDrawer := asci.NewAsci()
	match := game.NewMatch(pair.PlayerA, pair.PlayerB)
	match.StartMatch(asciDrawer)
}
