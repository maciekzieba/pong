package io

import "github.com/maciekzieba/pong/object"

type Controller interface {
	Move() object.RacketVectorValue
}
