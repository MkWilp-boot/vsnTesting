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
		Title:  "BHELL",
		Bounds: pixel.R(0, 0, 600, 650),
		VSync:  true,
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
			Speed:  100.0,
			Size:   playerSprite.Picture().Bounds(),
		},
		Moving: false,
	}

	last := time.Now()
	count := time.Now().Add(time.Millisecond * 100)

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
		for _, v := range entitys.PlayerFiredBullet {
			v.Sprite.Draw(win, pixel.IM.Moved(v.Pos))
			v.Tick(dt)
		}

		player.Sprite.Draw(win, pixel.IM.Moved(player.Pos))
		win.Update()
		player.Moving = false
	}
}

func main() {
	pixelgl.Run(run)
}
