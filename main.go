package main

import (
	"bhell/entitys"
	imageutils "bhell/imageUtils"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/exp/shiny/materialdesign/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:     "BHELL",
		Bounds:    pixel.R(0, 0, 600, 650),
		Resizable: false,
		Maximized: true,
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
			Pos:    win.Bounds().Center(),
			Speed:  200.0,
			Size:   playerSprite.Picture().Bounds(),
		},
		Moving: false,
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
		player.Moving = false
		win.Update()
	}
}

func updateBullets(win *pixelgl.Window, dt float64) {
	for _, v := range entitys.PlayerFiredBullet {
		v.Sprite.Draw(win, pixel.IM.Moved(v.Pos))
		v.Tick(win, dt)
	}
}

func main() {
	pixelgl.Run(run)
}
