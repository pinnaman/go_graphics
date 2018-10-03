package main

import (
	"fmt"
	"goray/raytrc"
	"image/color"
	"time"
)

func main() {
	var width = 400
	var height = 200
	scene := raytrc.NewScene(width, height)
	scene.EachPixel(func(x, y int) color.RGBA {
		return color.RGBA{
			uint8(x * 255 / width),
			uint8(y * 255 / height),
			100,
			255,
		}
	})
	scene.Save(fmt.Sprintf("./renders/%d.png", time.Now().Unix()))
}
