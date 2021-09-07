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
	Rect   pixel.Rect
	Life   int
}

type Enemy struct {
	Entity
}

type Bullet struct {
	Entity
}

func (bullet *Bullet) CheckCollisionWithEntity(enemy *Enemy, dt float64) {
	go func(bullet *Bullet, enemy *Enemy, dt float64) {
		enR := pixel.Rect{
			Min: pixel.V(bullet.Pos.X-(bullet.Size.W()/2), bullet.Pos.Y-(bullet.Size.H()/2)),
			Max: pixel.V(bullet.Pos.X+(bullet.Size.W()/2), bullet.Pos.Y+(bullet.Size.H()/2)),
		}
		coR := pixel.Rect{
			Min: pixel.V(enemy.Pos.X-(enemy.Size.W()/2), enemy.Pos.Y-(enemy.Size.H()/2)),
			Max: pixel.V(enemy.Pos.X+(enemy.Size.W()/2), enemy.Pos.Y+(enemy.Size.H()/2)),
		}
		if enR.Intersects(coR) {
			enemy.Life = enemy.Life - 1
			go removeBullet(bullet)
		}
	}(bullet, enemy, dt)
}

func removeBullet(bullet *Bullet) {
	for index, value := range PlayerFiredBullet {
		if value == bullet {
			PlayerFiredBullet = append(PlayerFiredBullet[:index], PlayerFiredBullet[index+1:]...)
		}
	}
}

func (bullet *Bullet) Tick(win *pixelgl.Window, dt float64) *Bullet {
	bullet.Pos.Y += bullet.Speed * dt
	if bullet.Pos.Y >= win.Bounds().Max.Y {
		removeBullet(bullet)
	}
	return bullet
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
	}
	bullet.Rect = pixel.Rect{
		Min: pixel.V(bullet.Pos.X-(bullet.Size.W()/2), bullet.Pos.Y-(bullet.Size.H()/2)),
		Max: pixel.V(bullet.Pos.X+(bullet.Size.W()/2), bullet.Pos.Y+(bullet.Size.H()/2)),
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
