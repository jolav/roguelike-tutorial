/* */

package main

import blt "roguelike/bearlibterminal"

type Tile struct {
	Blocked  bool
	BlockLOS bool
	Explored bool
	Visible  bool
}

func (t *Tile) render(x, y int) {
	blt.Layer(0)
	if t.Visible {
		if t.Blocked {
			//blt.BkColor(blt.ColorFromName("darkgb"))
			blt.Color(blt.ColorFromName("white"))
			blt.Print(x, y, "#")
		} else {
			//blt.BkColor(blt.ColorFromName("darkbg"))
			blt.Color(blt.ColorFromName("white"))
			blt.Print(x, y, ".")
		}
	} else if t.Explored {
		if t.Blocked {
			//blt.BkColor(blt.ColorFromName("darkgb"))
			blt.Color(blt.ColorFromName("gray"))
			blt.Print(x, y, "#")
		} else {
			//blt.BkColor(blt.ColorFromName("darkbg"))
			blt.Color(blt.ColorFromName("gray"))
			blt.Print(x, y, ".")
		}
	}

}
