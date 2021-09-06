package entitys

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Entity struct {
	Sprite *pixel.Sprite
	Pos    pixel.Vec
	Speed  float64
	Size   pixel.Rect
}

type Bullet struct {
	Entity
}

type Player struct {
	Entity
	Moving bool
	Firing bool
}

func (player *Player) MovementHandler(win *pixelgl.Window, dt float64) {
	// Player UP and Down
	if win.Pressed(pixelgl.KeyW) {
		player.Pos.Y += player.Speed * float64(dt)
	} else if win.Pressed(pixelgl.KeyS) {
		player.Pos.Y -= player.Speed * float64(dt)
	}
	//Player Left and Right
	if win.Pressed(pixelgl.KeyA) {
		player.Pos.X -= player.Speed * float64(dt)
	} else if win.Pressed(pixelgl.KeyD) {
		player.Pos.X += player.Speed * float64(dt)
	}
}
