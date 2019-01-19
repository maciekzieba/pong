package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardConstructor(t *testing.T) {
	board := NewBoard(10)
	assert.Equal(t, 30, board.Width)
	assert.Equal(t, 10, board.Height)

}
