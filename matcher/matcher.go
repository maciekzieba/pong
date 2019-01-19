package matcher

import (
	"github.com/maciekzieba/pong/player"
)

type (
	PlayerMatcher struct {
		players map[int]*player.Player
	}

	PlayerPair struct {
		PlayerA *player.Player
		PlayerB *player.Player
	}
)

func NewPlayerMatcher() *PlayerMatcher {
	return &PlayerMatcher{players: make(map[int]*player.Player)}
}

func (m *PlayerMatcher) AddPlayer(player *player.Player) {
	index := len(m.players)
	m.players[index] = player
}

func (m *PlayerMatcher) TryMatch() *PlayerPair {
	if len(m.players) > 1 {
		pair := &PlayerPair{}
		keys := make([]int, 0, len(m.players))
		for index := range m.players {
			keys = append(keys, index)
		}
		indexPlayerA := keys[0]
		indexPlayerB := keys[1]

		pair.PlayerA = m.players[indexPlayerA]
		delete(m.players, indexPlayerA)

		pair.PlayerB = m.players[indexPlayerB]
		delete(m.players, indexPlayerB)

		return pair
	}
	return nil
}
