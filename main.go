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
		},
	}

	last := time.Now()

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

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

		win.Clear(colornames.Black)
		player.Sprite.Draw(win, pixel.IM.Moved(player.Pos))
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
