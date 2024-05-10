/* */

package main

import (
	blt "roguelike/bearlibterminal"
	"strconv"
)

func renderAll(w *World) {
	blt.Clear()
	w.Camera.move(w.Entities["player"], w.Map)
	//fmt.Println(w.Entities["player"].X, w.Entities["player"].Y)
	renderMap(w.Map, w.Camera)
	renderEntities(w.Camera, w.Entities)
	//blt.PrintExt(5, 5, 200, 30, blt.TK_ALIGN_DEFAULT, "[color=orange][font=ubuntu]HI ROGUELIKE-DEV[/color]")
	blt.Refresh()
}

func renderEntities(cam *Camera, es Entities) {
	for _, e := range es {
		camX, camY := cam.mapToCameraCoords(e.X, e.Y)
		e.render(camX, camY)
	}
}

func renderMap(m *GameMap, cam *Camera) {
	for x := cam.X; x < cam.X+cam.Width; x++ {
		for y := cam.Y; y < cam.Y+cam.Height; y++ {
			camX, camY := cam.mapToCameraCoords(x, y)
			m.Tiles[x][y].render(camX, camY)
		}
	}
}

func initBLT(c Configuration) {
	blt.Open()

	fontSize2 := "size=" + strconv.Itoa(c.FontSize2)
	blt.Set("ubuntu font: " + c.Font2 + ", " + fontSize2)

	size := "size=" + strconv.Itoa(c.ScreenWidth) + "x"
	size += strconv.Itoa(c.ScreenHeight)
	title := "title='" + c.Title + "'"
	window := "window: " + size + "," + title
	fontSize := "size=" + strconv.Itoa(c.SquareSize)
	font := "font: " + c.Square + ", " + fontSize
	blt.Set(window + "; " + font)

	/*
		grey or gray, red, flame, orange, amber, yellow, lime, chartreuse, green, sea, turquoise, cyan, sky, azure, blue, han, violet, purple, fuchsia, magenta, pink, crimson, transparent
	*/
	blt.Set("palette.darkGrey=#636466")
	blt.Set("palette.choco=#bdb897")
	blt.Set("palette.purple=#dd00d7")
	blt.Set("palette.violet=#b897db")
	blt.Set("palette.khaki=#d1cdb6")
	blt.Set("palette.darkbg=#2f2f2f")
	blt.Set("palette.lightgrey=#444444")

	//blt.BkColor(blt.ColorFromName("darkbg"))
	blt.Clear()
}

func readKeyboard() (key int) {
	return blt.Read()
}

func closeWindow() {
	blt.Close()
}
