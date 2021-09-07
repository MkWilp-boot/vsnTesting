package main

import (
	"bhell/entitys"
	imageutils "bhell/imageUtils"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/exp/shiny/materialdesign/colornames"
)

var Boss entitys.Enemy

func run() {
	cfg := pixelgl.WindowConfig{
		Title:     "BHELL",
		Bounds:    pixel.R(0, 0, 600, 650),
		Resizable: false,
		//Maximized: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// creating the player
	playerSprite, err := imageutils.LoadPicture("assets/player.png")
	if err != nil {
		panic(err)
	}
	player := entitys.Player{
		Entity: entitys.Entity{
			Sprite: playerSprite,
			Pos:    win.Bounds().Center().Sub(pixel.V(0, 150)),
			Speed:  200.0,
			Size:   playerSprite.Picture().Bounds(),
		},
		Moving: false,
	}

	// creating the player
	bossprite, err := imageutils.LoadPicture("assets/enemy.png")
	if err != nil {
		panic(err)
	}
	Boss = entitys.Enemy{
		Entity: entitys.Entity{
			Sprite: bossprite,
			Pos:    win.Bounds().Center().Add(pixel.V(0, 150)),
			Speed:  0,
			Size:   bossprite.Picture().Bounds(),
		},
	}

	last := time.Now()
	count := time.Now().Add(time.Millisecond * 100)

	// main loop
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		if win.Pressed(pixelgl.MouseButton1) {
			if last.UnixMilli() >= count.UnixMilli() {
				count = time.Now().Add(time.Millisecond * 100)
				player.FireHandler(win, dt)
			}
		}

		player.MovementHandler(win, dt)
		win.Clear(colornames.Black)
		updateBullets(win, dt)

		player.Sprite.Draw(win, pixel.IM.Moved(player.Pos))
		Boss.Sprite.Draw(win, pixel.IM.Moved(Boss.Pos))
		player.Moving = false
		win.Update()
	}
}

func updateBullets(win *pixelgl.Window, dt float64) {
	for _, v := range entitys.PlayerFiredBullet {
		v.Sprite.Draw(win, pixel.IM.Moved(v.Pos))
		v.Tick(win, dt).CheckCollisionWithEntity(&Boss.Entity, dt)
	}
}

func main() {
	pixelgl.Run(run)
}
