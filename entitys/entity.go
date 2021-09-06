package entitys

import (
	imageutils "bhell/imageUtils"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var PlayerFiredBullet = make([]*Bullet, 0)

type Entity struct {
	Sprite *pixel.Sprite
	Pos    pixel.Vec
	Speed  float64
	Size   pixel.Rect
}

type Bullet struct {
	Entity
	Life    int
	MaxLife int
}

func (bullet *Bullet) Tick(dt float64) {
	bullet.Pos.Y += bullet.Speed * dt
	bullet.Life++
	if bullet.Life >= bullet.MaxLife {
		for i, v := range PlayerFiredBullet {
			if v == bullet {
				s := append(PlayerFiredBullet[:i], PlayerFiredBullet[i+1:]...)
				PlayerFiredBullet = s
			}
		}
	}
}

type Player struct {
	Entity
	Moving bool
}

func (player *Player) FireHandler(win *pixelgl.Window, dt float64) {
	bulletSprite, err := imageutils.LoadPicture("assets/bullet.png")
	if err != nil {
		panic(err)
	}

	bulletPos := player.Pos
	newPos := pixel.V(0, 10)
	bulletPos = bulletPos.Add(newPos)

	bullet := Bullet{
		Entity: Entity{
			Sprite: bulletSprite,
			Pos:    bulletPos,
			Speed:  500.0,
			Size:   bulletSprite.Picture().Bounds(),
		},
		Life:    0,
		MaxLife: 600,
	}
	PlayerFiredBullet = append(PlayerFiredBullet, &bullet)
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
