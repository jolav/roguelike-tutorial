/* */

package main

import (
	"math"
)

type FieldOfVision struct {
	sin       map[int]float64
	cos       map[int]float64
	losRadius int
}

func (f *FieldOfVision) initFOV() {
	f.cos = make(map[int]float64)
	f.sin = make(map[int]float64)

	for i := 0; i < 360; i++ {
		ax := math.Sin(float64(i) / (float64(180) / math.Pi))
		ay := math.Cos(float64(i) / (float64(180) / math.Pi))
		f.sin[i] = ax
		f.cos[i] = ay
	}
	f.losRadius = 5
}

func (f *FieldOfVision) rayCast(w *World) {
	// clean map
	for x := 0; x < w.Map.Width; x++ {
		for y := 0; y < w.Map.Height; y++ {
			w.Map.Tiles[x][y].Visible = false
			//blt.Print(x, y, " ")
		}
	}
	// mark player
	pX := w.Entities["player"].X
	pY := w.Entities["player"].Y
	w.Map.Tiles[pX][pY].Explored = true
	w.Map.Tiles[pX][pY].Visible = true

	for i := 0; i < 360; i++ {
		x := float64(pX)
		y := float64(pY)

		for r := 0; r < w.Fov.losRadius; r++ {
			x += f.sin[i]
			y += f.cos[i]
			roundedX := int(round(x))
			roundedY := int(round(y))
			if x < 0 || x > float64(w.Map.Width-1) || y < 0 || y > float64(w.Map.Height-1) {
				break
			}
			w.Map.Tiles[roundedX][roundedY].Explored = true
			w.Map.Tiles[roundedX][roundedY].Visible = true
			if w.Map.Tiles[roundedX][roundedY].BlockLOS == true {
				break
			}
		}
	}
}
