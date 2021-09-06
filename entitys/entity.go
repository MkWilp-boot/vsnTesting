package entitys

import "github.com/faiface/pixel"

type Entity struct {
	Sprite *pixel.Sprite
	Pos    pixel.Vec
	Speed  float64
}

type Player struct {
	Entity
}
