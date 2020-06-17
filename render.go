/* */

package main

import (
	blt "roguelike/bearlibterminal"
)

func renderAll(w *World) {
	blt.Clear()
	renderMap(w.Map)
	for _, e := range w.Entities {
		renderEntity(e)
	}
	blt.Refresh()
}

func renderMap(m *GameMap) {
	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			renderTile(x, y, m.Tiles[x][y])
		}
	}
}

func renderTile(x, y int, t *Tile) {
	blt.Layer(0)
	if t.Blocked {
		blt.BkColor(blt.ColorFromName("darkGrey"))
		blt.Color(blt.ColorFromName("choco"))
		blt.Print(x, y, "#")
	} else {
		blt.BkColor(blt.ColorFromName("khaki"))
		blt.Color(blt.ColorFromName("violet"))
		blt.Print(x, y, ".")
	}
}

func renderEntity(e *Entity) {
	blt.Layer(0)
	blt.BkColor(blt.ColorFromName("khaki"))
	blt.Color(blt.ColorFromName(e.Color))
	blt.Print(e.X, e.Y, e.Char)
}
