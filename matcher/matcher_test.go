package matcher

import (
	"github.com/maciekzieba/pong/player"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldNotMatchOnePlayer(t *testing.T) {
	p1 := player.NewPlayer("1")
	matcher := NewPlayerMatcher()
	matcher.AddPlayer(p1)
	pair := matcher.TryMatch()
	assert.Nil(t, pair)
}

func TestShouldMatchTwoPlayers(t *testing.T) {
	p1 := player.NewPlayer("1")
	p2 := player.NewPlayer("2")

	matcher := NewPlayerMatcher()
	matcher.AddPlayer(p1)
	matcher.AddPlayer(p2)

	pair := matcher.TryMatch()
	assert.NotNil(t, pair)
	assert.Equal(t, "1", pair.PlayerA.Name)
	assert.Equal(t, "2", pair.PlayerB.Name)
}
