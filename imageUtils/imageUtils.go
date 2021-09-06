package imageutils

import (
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
)

func toSprite(pic pixel.Picture) *pixel.Sprite {
	return pixel.NewSprite(pic, pic.Bounds())
}

func LoadPicture(path string) (*pixel.Sprite, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	sprite := toSprite(pixel.PictureDataFromImage(img))
	return sprite, nil
}
