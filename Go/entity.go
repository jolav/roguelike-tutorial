/* */

package main

import blt "roguelike/bearlibterminal"

type Entities map[string]*Entity

type Entity struct {
	X      int
	Y      int
	Layer  int
	Char   string
	Color  string
	Facing rune
}

func (e *Entity) move(dx, dy int) {
	e.X += dx
	e.Y += dy
}

func (e *Entity) render(x, y int) {
	blt.Layer(e.Layer)
	//blt.BkColor(blt.ColorFromName("darkbg"))
	blt.Color(blt.ColorFromName(e.Color))
	blt.Print(x, y, e.Char)
}
