// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main is a simple demo to show how to use the hilbert library
package main

import (
	"flag"
	"image"
	"image/color"
	"image/gif"
	"log"
	"os"

	"math"
	"path/filepath"
	"strconv"

	"github.com/bramp/hilbert"
	"github.com/bramp/hilbert/demo/lib"
	"github.com/fogleman/gg"
)

// spaceFillingImage facilitates the drawing of a space filing curve.
type spaceFillingImage struct {
	Curve hilbert.SpaceFilling2D

	// Size of each square in pixels
	SquareWidth  float64
	SquareHeight float64

	DrawGrid   bool
	DrawText   bool    // Should text be drawn on the image
	TextMargin float64 // Margin around text in pixels

	BackgroundColor color.Color
	GridColor       color.Color
	TextColor       color.Color
	SnakeColor      color.Color

	GridWidth  float64
	SnakeWidth float64
}

// createSpaceFillingImage returns a new SpaceFillingImage ready for drawing.
func createSpaceFillingImage(curve hilbert.SpaceFilling2D, squareWidth, squareHeight float64) *spaceFillingImage {
	return &spaceFillingImage{
		Curve: curve,

		SquareWidth:  squareWidth,
		SquareHeight: squareHeight,

		DrawGrid:   true,
		DrawText:   true,
		TextMargin: 3.0,

		BackgroundColor: color.RGBA{0xee, 0xee, 0xff, 0xff},
		GridColor:       color.White,
		TextColor:       color.RGBA{0x33, 0x33, 0x33, 0xff},
		SnakeColor:      color.RGBA{0x33, 0x33, 0x33, 0xff},

		GridWidth:  1.0,
		SnakeWidth: 2.0,
	}
}

func (h *spaceFillingImage) toPixel(x, y int) (float64, float64) {
	switch h.Curve.GetGridType() {
	case hilbert.GridTriangular:
		triSide := h.SquareWidth
		triHeight := triSide * math.Sqrt(3) / 2
		width, _ := h.Curve.GetDimensions()
		totalWidth := float64(width) * triSide
		rowY := float64(y) * triHeight
		rowOffset := (totalWidth - float64(y+1)*triSide) / 2
		return rowOffset + float64(x)*triSide/2, rowY

	case hilbert.GridHexagonal:
		size := h.SquareWidth / 2
		px := size * 3 / 2 * float64(x)
		py := size * math.Sqrt(3) * (float64(y) + float64(x)/2)
		return px + 512, py + 512

	default:
		return float64(x) * h.SquareWidth, float64(y) * h.SquareHeight
	}
}

func (h *spaceFillingImage) drawGrid(gc *gg.Context, width, height int) {
	if h.Curve.GetGridType() != hilbert.GridSquare {
		return
	}
	for x := 0; x <= width; x++ {
		gc.MoveTo(h.toPixel(x, 0))
		gc.LineTo(h.toPixel(x, height))
	}
	for y := 0; y <= height; y++ {
		gc.MoveTo(h.toPixel(0, y))
		gc.LineTo(h.toPixel(width, y))
	}
	gc.SetLineWidth(h.GridWidth)
	gc.SetColor(h.GridColor)
	gc.Stroke()
}

func (h *spaceFillingImage) Draw() (*gg.Context, error) {
	width, height := h.Curve.GetDimensions()
	gridType := h.Curve.GetGridType()

	var pwidth, pheight float64
	switch gridType {
	case hilbert.GridTriangular:
		triSide := h.SquareWidth
		triHeight := triSide * math.Sqrt(3) / 2
		pwidth = float64(width) * triSide
		pheight = float64(height) * triHeight
	case hilbert.GridHexagonal:
		pwidth, pheight = 1024, 1024
	default:
		pwidth, pheight = float64(width)*h.SquareWidth, float64(height)*h.SquareHeight
	}

	gc := gg.NewContext(int(pwidth), int(pheight))
	gc.SetColor(h.BackgroundColor)
	gc.Clear()

	if h.DrawGrid {
		h.drawGrid(gc, width, height)
	}

	n := h.Curve.GetCount()
	for t := 0; t < n; t++ {
		x, y, err := h.Curve.Map(t)
		if err != nil {
			return nil, err
		}

		px, py := h.toPixel(x, y)
		if h.DrawText && gridType == hilbert.GridSquare {
			gc.SetColor(h.TextColor)
			gc.DrawStringAnchored(strconv.Itoa(t), px+h.TextMargin, py, 0, 1)
		}

		var cx, cy float64
		switch gridType {
		case hilbert.GridTriangular:
			triSide := h.SquareWidth
			triHeight := triSide * math.Sqrt(3) / 2
			if x%2 == 0 {
				cx, cy = px+triSide/2, py+triHeight*2/3
			} else {
				cx, cy = px+triSide/2, py+triHeight/3
			}
		case hilbert.GridHexagonal:
			cx, cy = px, py
		default:
			cx, cy = px+h.SquareWidth/2, py+h.SquareHeight/2
		}

		if t == 0 {
			gc.MoveTo(cx, cy)
		} else {
			gc.LineTo(cx, cy)
		}
	}

	gc.SetColor(h.SnakeColor)
	gc.SetLineWidth(h.SnakeWidth)
	gc.SetLineCap(gg.LineCapSquare)
	gc.SetLineJoin(gg.LineJoinRound)
	gc.Stroke()

	return gc, nil
}

func mainDrawOne(filename string, curve hilbert.SpaceFilling2D) error {
	log.Printf("Drawing one image %q", filename)
	img, err := createSpaceFillingImage(curve, 64, 64).Draw()
	if err != nil {
		return err
	}
	return img.SavePNG(filename)
}

func mainDrawAnimation(filename string, factory func(n int) hilbert.SpaceFilling2D, min, max int) error {
	log.Printf("Drawing animation %q", filename)
	iterations := max - min
	g := gif.GIF{
		Image:     make([]*image.Paletted, iterations),
		Delay:     make([]int, iterations),
		LoopCount: 0,
	}

	const canvasSize = 512.0
	for i := 0; i < iterations; i++ {
		curve := factory(min + i)
		w, h := curve.GetDimensions()
		
		gridType := curve.GetGridType()
		var sw, sh float64
		if gridType == hilbert.GridTriangular {
			sw = canvasSize / float64(w)
			sh = sw // triangular height will be calculated from this
		} else if gridType == hilbert.GridHexagonal {
			sw = 32.0 // Fixed size for hex
			sh = 32.0
		} else {
			sw = canvasSize / float64(w)
			sh = canvasSize / float64(h)
		}

		himg := createSpaceFillingImage(curve, sw, sh)
		himg.DrawText = false
		img, err := himg.Draw()
		if err != nil {
			return err
		}
		g.Image[i] = lib.ConvertToPaletted(img.Image())
		g.Delay[i] = 100
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	return gif.EncodeAll(f, &g)
}

func mainDrawLogo(filename string, curve hilbert.SpaceFilling2D) error {
	log.Printf("Drawing logo %q", filename)
	h := createSpaceFillingImage(curve, 256, 256)
	h.DrawText, h.DrawGrid = false, false
	h.SnakeWidth = 64
	h.BackgroundColor = color.Transparent
	img, err := h.Draw()
	if err != nil {
		return err
	}
	return img.SavePNG(filename)
}

func main() {
	algo := flag.String("algo", "all", "Algorithm")
	output := flag.String("output", "", "Output filename")
	logo := flag.Bool("logo", false, "Draw logo version")
	flag.Parse()

	factories := map[string]func(n int) hilbert.SpaceFilling2D{
		"hilbert": func(n int) hilbert.SpaceFilling2D {
			s, _ := hilbert.NewHilbert(int(math.Pow(2, float64(n))))
			return s
		},
		"peano": func(n int) hilbert.SpaceFilling2D {
			s, _ := hilbert.NewPeano(int(math.Pow(3, float64(n))))
			return s
		},
		"morton": func(n int) hilbert.SpaceFilling2D {
			s, _ := hilbert.NewMorton(int(math.Pow(2, float64(n))))
			return s
		},
		"moore": func(n int) hilbert.SpaceFilling2D {
			s, _ := hilbert.NewMoore(int(math.Pow(2, float64(n))))
			return s
		},
		"sierpinski": func(n int) hilbert.SpaceFilling2D {
			s, _ := hilbert.NewSierpinski(int(math.Pow(2, float64(n))))
			return s
		},
		"snake": func(n int) hilbert.SpaceFilling2D {
			s, _ := hilbert.NewSnake(int(math.Pow(2, float64(n))))
			return s
		},
		"gosper": func(n int) hilbert.SpaceFilling2D {
			s, _ := hilbert.NewGosper(n)
			return s
		},
	}

	draw := func(name string, out string) {
		factory := factories[name]
		var err error
		if *logo {
			err = mainDrawLogo(out, factory(4))
		} else if filepath.Ext(out) == ".gif" {
			err = mainDrawAnimation(out, factory, 1, 5)
		} else {
			n := 3
			if name == "peano" { n = 2 }
			if name == "gosper" { n = 2 }
			err = mainDrawOne(out, factory(n))
		}
		if err != nil {
			log.Fatalf("Failed to draw %s: %v", name, err)
		}
	}

	if *algo == "all" {
		for name := range factories {
			draw(name, name+".png")
			draw(name, name+"_animation.gif")
		}
	} else {
		out := *output
		if out == "" { out = *algo + ".png" }
		draw(*algo, out)
	}
}
