// Package animated_gif creates gif mathematically
package animated_gif

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand/v2"
)

type LissajousFigure struct {
	Palette                              []color.Color
	Width, Height, Frames, Delay, Cycles int
	Resolution                           float64
}

// DrawLissajousFigure writes a LissajousFigure gif to an io.Writer
// example:
// f, err := os.Create("new_file.gif")
// DrawLissajousFigure(f, LissajousFigure{})
func DrawLissajousFigure(
	output io.Writer,
	l LissajousFigure,
) (err error) {
	var palette []color.Color
	var width, height, frames, cycles, delay int
	var res float64
	if l.Palette == nil || len(l.Palette) == 0 {
		for range 2 {
			palette = append(palette, color.RGBA{
				R: uint8(rand.UintN(256)),
				G: uint8(rand.UintN(256)),
				B: uint8(rand.UintN(256)),
				A: uint8(rand.UintN(256)),
			})
		}
	} else if len(l.Palette) == 1 {
		palette = append(l.Palette, color.RGBA{
			R: uint8(rand.UintN(256)),
			G: uint8(rand.UintN(256)),
			B: uint8(rand.UintN(256)),
			A: uint8(rand.UintN(256)),
		})
	} else {
		palette = l.Palette
	}
	if l.Width == 0 {
		width = 256
	} else {
		width = l.Width
	}
	if l.Height == 0 {
		height = 256
	} else {
		height = l.Height
	}
	if l.Frames == 0 {
		frames = 64
	} else {
		frames = l.Frames
	}
	if l.Delay == 0 {
		delay = 8
	} else {
		delay = l.Delay / 10
	}
	if l.Cycles == 0 {
		cycles = 10
	} else {
		cycles = l.Cycles
	}
	if l.Resolution == 0 {
		res = 0.001
	} else {
		res = l.Resolution
	}
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: frames}
	phase := 0.0
	for range frames {
		rect := image.Rect(0, 0, width, height)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				width/2+int(float64(width)/2*x+0.5), height/2+int(float64(height)/2*y+0.5),
				uint8(1+rand.UintN(uint(len(palette)-1))),
			)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err = gif.EncodeAll(output, &anim)
	return
}
