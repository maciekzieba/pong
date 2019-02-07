package keyboard

import (
	"fmt"
	"math/rand"

	"github.com/maciekzieba/pong/object"
)

type KeyboardController struct {
}

func NewKeyboardController() *KeyboardController {
	controller := &KeyboardController{}
	return controller
}

func (k *KeyboardController) Move() object.RacketVectorValue {
	//rand.Seed(time.Now().Unix())
	v := rand.Intn(10) + 1
	fmt.Printf(" v: %d", v)
	switch v {
	case 2:
		return object.RacketVectorDown
	case 3:
		return object.RacketVectorUp
	}
	return object.RacketVectorNone
}
