/* */

package main

import blt "roguelike/bearlibterminal"

type Tile struct {
	Blocked  bool
	BlockLOS bool
}

func (t *Tile) render(x, y int) {
	blt.Layer(0)
	if t.Blocked {
		//blt.BkColor(blt.ColorFromName("darkgb"))
		blt.Color(blt.ColorFromName("violet"))
		blt.Print(x, y, "#")
	} else {
		//blt.BkColor(blt.ColorFromName("darkbg"))
		blt.Color(blt.ColorFromName("violet"))
		blt.Print(x, y, ".")
	}
}
