package main

import (
	"examples/ch1/animated_gif"
	"fmt"
	"image/color"
	"os"
)

func main() {
	var path string
	if len(os.Args) >= 2 {
		path = os.Args[1]
	}
	if path == "" {
		path = "animated_gif.gif"
	}
	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("error creating file: %+v", err)
	}
	err = animated_gif.DrawLissajousFigure(
		f,
		animated_gif.LissajousFigure{
			Palette: []color.Color{
				color.RGBA{R: 255, G: 255, B: 255, A: 255},
				color.RGBA{R: 255, G: 0, B: 0, A: 255},
				color.RGBA{R: 0, G: 255, B: 0, A: 255},
				color.RGBA{R: 0, G: 0, B: 255, A: 255},
			},
			Width:      1920,
			Height:     1080,
			Frames:     100,
			Delay:      100,
			Cycles:     15,
			Resolution: 0.001,
		},
	)
	if err != nil {
		fmt.Printf("error creating animated gif: %+v", err)
	} else {
		fmt.Printf("GIF has been created at %s", path)
	}
}
